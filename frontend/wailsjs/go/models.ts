export namespace backend {
	
	export class KeyInfo {
	    key: string;
	    description: string;
	    type: string;
	    group: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.description = source["description"];
	        this.type = source["type"];
	        this.group = source["group"];
	    }
	}
	export class GameSchema {
	    id: string;
	    name: string;
	    fileName: string;
	    section: string;
	    keys: KeyInfo[];
	
	    static createFrom(source: any = {}) {
	        return new GameSchema(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.fileName = source["fileName"];
	        this.section = source["section"];
	        this.keys = this.convertValues(source["keys"], KeyInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

