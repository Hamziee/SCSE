import './style.css';
import { GetGames, LoadSave, SaveData } from '../wailsjs/go/main/App';

let currentGame = null;
let currentData = {};
let currentSchema = null;

const setup = async () => {
    try {
        const games = await GetGames();
        renderSidebar(games);
        setTimeout(() => {
            const splash = document.getElementById('splash-screen');
            if (splash) splash.classList.add('hidden');
        }, 800);
    } catch (e) {
        console.error("Failed to load games:", e);
    }

    document.getElementById('btn-save').addEventListener('click', saveChanges);
    document.getElementById('btn-reload').addEventListener('click', () => {
        if (currentGame) {
            loadGame(currentGame, currentSchema);
        }
    });
};

function renderSidebar(games) {
    const list = document.getElementById('game-list');
    list.innerHTML = '';

    if (!games || games.length === 0) {
        list.innerHTML = '<div style="padding:15px; color:var(--text-color)">No games definitions found.</div>';
        return;
    }

    games.forEach(game => {
        const div = document.createElement('div');
        div.className = 'game-item';
        div.textContent = game.name;
        div.onclick = () => {
            document.querySelectorAll('.game-item').forEach(el => el.classList.remove('active'));
            div.classList.add('active');
            loadGame(game.id, game);
        };
        list.appendChild(div);
    });
}

async function loadGame(gameId, schema) {
    currentGame = gameId;
    currentSchema = schema;
    document.getElementById('current-game-title').textContent = schema.name;
    document.getElementById('actions').style.display = 'flex';
    document.getElementById('content-area').innerHTML = '<p style="text-align:center; margin-top:50px;">Loading save file...</p>';

    try {
        const data = await LoadSave(gameId);
        currentData = data || {};
        renderEditor(currentData, schema);
    } catch (e) {
        document.getElementById('content-area').innerHTML = `<div style="padding:20px; text-align:center;"><p style="color:#ff6b6b; font-weight:bold;">Error loading save file</p><p>${e}</p><p style="font-size:0.9em; opacity:0.8; margin-top:10px;">Make sure the game is installed or the save file exists in AppData/Roaming/MMFApplications.</p></div>`;
        console.error(e);
    }
}

function renderEditor(data, schema) {
    const container = document.getElementById('content-area');
    container.innerHTML = '<div id="editor-grid" class="fade-in-up"></div>';
    const grid = document.getElementById('editor-grid');

    let currentGroup = null;
    let groupContainer = null;
    let groupGrid = null;

    schema.keys.forEach(keyInfo => {

        if (keyInfo.group !== currentGroup) {
            currentGroup = keyInfo.group;

            if (currentGroup) {

                groupContainer = document.createElement('details');
                groupContainer.className = 'group-container';
                groupContainer.style.gridColumn = '1 / -1';
                groupContainer.style.width = '100%';
                groupContainer.style.backgroundColor = 'rgba(255, 255, 255, 0.03)';
                groupContainer.style.borderRadius = '12px';
                groupContainer.style.overflow = 'hidden';

                const summary = document.createElement('summary');
                summary.textContent = currentGroup;
                summary.style.padding = '15px';
                summary.style.cursor = 'pointer';
                summary.style.fontWeight = 'bold';
                summary.style.color = 'var(--heading-color)';
                summary.style.display = 'list-item';

                groupContainer.appendChild(summary);


                groupGrid = document.createElement('div');
                groupGrid.style.display = 'grid';
                groupGrid.style.gridTemplateColumns = 'repeat(auto-fill, minmax(250px, 1fr))';
                groupGrid.style.gap = '20px';
                groupGrid.style.padding = '20px';

                groupGrid.style.marginTop = '-10px';
                groupContainer.appendChild(groupGrid);
                grid.appendChild(groupContainer);
            } else {
                groupContainer = null;
                groupGrid = null;
            }
        }

        const card = document.createElement('div');
        card.className = 'field-card';

        const fullKey = `${schema.section}|${keyInfo.key}`;
        const val = data[fullKey];

        const label = document.createElement('label');
        label.className = 'field-label';
        label.textContent = keyInfo.description || keyInfo.key;

        let input;
        if (keyInfo.type === 'bool') {
            const wrapper = document.createElement('div');
            wrapper.className = 'checkbox-wrapper';

            const txt = document.createElement('span');
            txt.textContent = keyInfo.description || keyInfo.key;
            txt.style.color = 'var(--heading-color)';

            input = document.createElement('input');
            input.type = 'checkbox';
            input.checked = (val === '1');

            input.onchange = (e) => {
                data[fullKey] = e.target.checked ? '1' : '0';
            };

            wrapper.appendChild(txt);
            wrapper.appendChild(input);
            card.appendChild(wrapper);
        } else {
            card.appendChild(label);
            input = document.createElement('input');
            input.className = 'field-input';

            input.value = val !== undefined ? val : '';
            if (keyInfo.type === 'int') {
                input.type = 'number';
            } else {
                input.type = 'text';
            }

            input.onchange = (e) => {
                data[fullKey] = e.target.value;
            };
            card.appendChild(input);
        }

        if (groupGrid) {
            groupGrid.appendChild(card);
        } else {
            grid.appendChild(card);
        }
    });

    if (schema.keys.length === 0) {
        container.innerHTML += '<p style="text-align:center; margin-top:20px;">No specific editor controls defined for this game yet.</p>';
    }
}

async function saveChanges() {
    if (!currentGame) return;
    try {
        await SaveData(currentGame, currentData);
        showToast();
    } catch (e) {
        alert("Error saving: " + e);
    }
}

function showToast() {
    const t = document.getElementById('toast');
    t.classList.add('show');
    setTimeout(() => t.classList.remove('show'), 3000);
}

setup();
