export namespace main {
	
	export class DirEntry {
	    Path: string;
	    Name: string;
	    IsDir: boolean;
	    Content: string;
	
	    static createFrom(source: any = {}) {
	        return new DirEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Name = source["Name"];
	        this.IsDir = source["IsDir"];
	        this.Content = source["Content"];
	    }
	}
	export class ReturnValue {
	    DirEntries: DirEntry[];
	    SelectedNote: DirEntry;
	    Bool: boolean;
	    Error: string;
	
	    static createFrom(source: any = {}) {
	        return new ReturnValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DirEntries = this.convertValues(source["DirEntries"], DirEntry);
	        this.SelectedNote = this.convertValues(source["SelectedNote"], DirEntry);
	        this.Bool = source["Bool"];
	        this.Error = source["Error"];
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

