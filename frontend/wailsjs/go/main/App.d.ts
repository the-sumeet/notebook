// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {context} from '../models';

export function DeleteNoteFile(arg1:string):Promise<main.ReturnValue>;

export function GetFiles():Promise<main.ReturnValue>;

export function GetSelectedNote():Promise<main.DirEntry>;

export function GoBack():Promise<main.ReturnValue>;

export function Greet(arg1:string):Promise<string>;

export function IsNotesDir():Promise<main.ReturnValue>;

export function NewFolder(arg1:string):Promise<main.ReturnValue>;

export function NewNote(arg1:string):Promise<main.ReturnValue>;

export function SaveFileContent(arg1:string,arg2:string):Promise<main.ReturnValue>;

export function SelectDir(arg1:string):Promise<main.ReturnValue>;

export function SelectNote(arg1:string):Promise<main.ReturnValue>;

export function SetTitle(arg1:context.Context,arg2:string):Promise<void>;
