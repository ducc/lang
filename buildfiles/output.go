// auto generated - changes will not be persisted
package main
import (
"fmt"
"github.com/ducc/lang/util"
"github.com/ducc/lang/builtins"
"go.uber.org/zap"
"os"
"github.com/bwmarrin/discordgo"
)
// incase imports are not used
var _ = fmt.Println
var _ = util.NewScope
var _ = builtins.Add
var _ = zap.S
// TODO IMPORTS
func main() {
logger, _ := zap.NewDevelopment()
zap.ReplaceGlobals(logger)
defer logger.Sync()
scope := util.NewScope()
{
scope := scope.Clone()
scope = scope.PushStack("STARTING")
// Node=CallFunction{Name: Log} Last=CallFunction{Name: processMessages} Equal=false Main=true
{
scope := scope.Clone()
zap.S().Debug("call                  Log ", scope.String())
scope = Log(scope)
}
}
{
scope := scope.Clone()
zap.S().Debug("call             getToken ", scope.String())
scope = getToken(scope)
// Node=CallFunction{Name: NewDiscord} Last=CallFunction{Name: processMessages} Equal=false Main=true
{
scope := scope.Clone()
zap.S().Debug("call           NewDiscord ", scope.String())
scope = NewDiscord(scope)
// Node=CallFunction{Name: processMessages} Last=CallFunction{Name: processMessages} Equal=true Main=true
{
scope := scope.Clone()
zap.S().Debug("call      processMessages ", scope.String())
scope = processMessages(scope)
}
}
}
}

func getToken(scope *util.Scope) *util.Scope {
{
scope := scope.Clone()
scope = scope.PushStack("DISCORD_TOKEN")
// Node=CallFunction{Name: GetEnv} Last=CallFunction{Name: GetEnv} Equal=true Main=false
{
scope := scope.Clone()
zap.S().Debug("call               GetEnv ", scope.String())
scope = GetEnv(scope)
return scope
}
}
}

func processMessages(scope *util.Scope) *util.Scope {
{
scope := scope.Clone()
/////// MEME ////////
scope = scope.GetVar("nextMessage").(func(*util.Scope) *util.Scope)(scope)
// Node=CallFunction{Name: ExtractMessageAttrs} Last=CallFunction{Name: processMessages} Equal=false Main=false
{
scope := scope.Clone()
zap.S().Debug("call  ExtractMessageAttrs ", scope.String())
scope = ExtractMessageAttrs(scope)
// Node=CallFunction{Name: handleCommand} Last=CallFunction{Name: processMessages} Equal=false Main=false
{
scope := scope.Clone()
zap.S().Debug("call        handleCommand ", scope.String())
scope = handleCommand(scope)
}
}
}
{
scope := scope.Clone()
zap.S().Debug("call      processMessages ", scope.String())
scope = processMessages(scope)
return scope
}
}

func handleCommand(scope *util.Scope) *util.Scope {
{
scope := scope.Clone()
scope = scope.PushStack("!help")
// Node=CallFunction{Name: EQ} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
zap.S().Debug("call                   EQ ", scope.String())
scope = EQ(scope)
// Node=DefineVariable{Name: sendHelp, Value: DefineFunctionValue{Name: sendHelp}} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
scope = scope.SetVar("sendHelp", sendHelp)
// Node=PushToStack{Variable: sendHelp} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
scope = scope.PushStack(scope.GetVar("sendHelp"))
// Node=CallFunction{Name: IfTrue} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
zap.S().Debug("call               IfTrue ", scope.String())
scope = IfTrue(scope)
}
}
}
}
}
{
scope := scope.Clone()
scope = scope.PushStack("!ping")
// Node=CallFunction{Name: EQ} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
zap.S().Debug("call                   EQ ", scope.String())
scope = EQ(scope)
// Node=DefineVariable{Name: sendPing, Value: DefineFunctionValue{Name: sendPing}} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
scope = scope.SetVar("sendPing", sendPing)
// Node=PushToStack{Variable: sendPing} Last=CallFunction{Name: IfTrue} Equal=false Main=false
{
scope := scope.Clone()
scope = scope.PushStack(scope.GetVar("sendPing"))
// Node=CallFunction{Name: IfTrue} Last=CallFunction{Name: IfTrue} Equal=true Main=false
{
scope := scope.Clone()
zap.S().Debug("call               IfTrue ", scope.String())
scope = IfTrue(scope)
return scope
}
}
}
}
}
}

func sendHelp(scope *util.Scope) *util.Scope {
{
scope := scope.Clone()
scope = scope.PushStack("Commands: !help, !ping")
// Node=CallFunction{Name: SendMessage} Last=CallFunction{Name: SendMessage} Equal=true Main=false
{
scope := scope.Clone()
zap.S().Debug("call          SendMessage ", scope.String())
scope = SendMessage(scope)
return scope
}
}
}

func sendPing(scope *util.Scope) *util.Scope {
{
scope := scope.Clone()
scope = scope.PushStack("Pong!")
// Node=CallFunction{Name: SendMessage} Last=CallFunction{Name: SendMessage} Equal=true Main=false
{
scope := scope.Clone()
zap.S().Debug("call          SendMessage ", scope.String())
scope = SendMessage(scope)
return scope
}
}
}

func GetEnv(scope *util.Scope) *util.Scope {
val, scope := scope.PopStack()
envvar := os.Getenv(val.(string))
return scope.PushStack(envvar)
}
func Log(scope *util.Scope) *util.Scope {
val, scope := scope.PopStack()
zap.S().Info(val)
return scope
}
func NewDiscord(scope *util.Scope) *util.Scope {
token, scope := scope.PopStack()
dg, err := discordgo.New("Bot " + token.(string))
if err != nil {
panic(err)
}
queue := make(chan *discordgo.MessageCreate)
dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
queue <- m
})
dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
err = dg.Open()
if err != nil {
panic(err)
}
scope = scope.SetVar("discord", dg)
scope = scope.SetVar("nextMessage", func(scope *util.Scope) *util.Scope {
msg := <-queue
return scope.PushStack(msg)
})
return scope
}
func ExtractMessageAttrs(scope *util.Scope) *util.Scope {
value, scope := scope.PopStack()
mc := value.(*discordgo.MessageCreate)
content := mc.Content
scope = scope.SetVar("channelID", mc.ChannelID)
return scope.PushStack(content)
}
func SendMessage(scope *util.Scope) *util.Scope {
cVal, scope := scope.PopStack()
content := cVal.(string)
channelID := scope.GetVar("channelID").(string)
discord := scope.GetVar("discord").(*discordgo.Session)
zap.S().Info(channelID, content)
_, err := discord.ChannelMessageSend(channelID, content)
if err != nil { panic(err) }
return scope
}
func EQ(scope *util.Scope) *util.Scope {
bVal, scope := scope.PopStack()
aVal, scope := scope.PopStack()
return scope.PushStack(aVal == bVal)
}
func IfTrue(scope *util.Scope) *util.Scope {
fVal, scope := scope.PopStack()
fun := fVal.(func(*util.Scope) *util.Scope)
bVal, scope := scope.PopStack()
boo := bVal.(bool)
if boo {
fun(scope)
}
return scope
}
func IfFalse(scope *util.Scope) *util.Scope {
fVal, scope := scope.PopStack()
fun := fVal.(func(*util.Scope) *util.Scope)
bVal, scope := scope.PopStack()
boo := bVal.(bool)
if !boo {
fun(scope)
}
return scope
}
