//go:build js && wasm

package main

import (
  "errors"
  "path/filepath"
  "syscall/js"

  "Peperouny/internal/modules"
  _ "Peperouny/internal/modules/impl"
)

var badCall = errors.New("usage: peperounyInspect/Decrypt(fileName, Uint8Array[, password])")

func main() {
  g:=js.Global()
  g.Set("peperounyInspect",js.FuncOf(inspect))
  g.Set("peperounyDecrypt",js.FuncOf(decrypt))
  select{}
}
func bytesFrom(v js.Value) []byte {
  if v.Type()!=js.TypeObject || v.Get("constructor").Get("name").String()!="Uint8Array" { return nil }
  b:=make([]byte,v.Get("length").Int()); js.CopyBytesToGo(b,v); return b
}
func inspect(_ js.Value,args []js.Value) any {
  if len(args)<2 { return fail(badCall) }
  name:=args[0].String(); data:=bytesFrom(args[1])
  mod:=modules.Find(filepath.Ext(name),"")
  if mod==nil { return fail(errors.New("unsupported file extension; use .npvt or .npvs")) }
  proto:=""
  if len(data)>=5 && string(data[:4])=="NPVS" { proto="NPVS" }
  if len(data)>=5 && string(data[:5])=="NPVT1" { proto="NPVT1" }
  needs:=false
  if mod.NeedsPassword!=nil { needs=mod.NeedsPassword(proto,string(data)) }
  format:="NPVT"
  if filepath.Ext(name)==".npvs" { format="NPVS" }
  return map[string]any{"ok":true,"format":format,"bytes":len(data),"module":mod.Name,"needsPassword":needs}
}
func decrypt(_ js.Value,args []js.Value) any {
  if len(args)<3 { return fail(badCall) }
  name:=args[0].String(); data:=bytesFrom(args[1]); password:=args[2].String()
  mod:=modules.Find(filepath.Ext(name),"")
  if mod==nil { return fail(errors.New("unsupported file extension; use .npvt or .npvs")) }
  proto:=""
  if len(data)>=5 && string(data[:4])=="NPVS" { proto="NPVS" }
  if len(data)>=5 && string(data[:5])=="NPVT1" { proto="NPVT1" }
  result,err:=mod.Decrypt(modules.Request{FileName:name,Data:data,Proto:proto,Payload:string(data),Password:password})
  if err!=nil { return fail(err) }
  return map[string]any{"ok":true,"text":result.Text,"fileName":result.FileName,"format":filepath.Ext(name)}
}
func fail(err error) map[string]any { return map[string]any{"ok":false,"error":err.Error()} }
