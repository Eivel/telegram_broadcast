os.loadAPI("json")

local box = peripheral.find("chatBox")
if box == nil then error("chatBox not found") end

print("Provide bot token")
local botToken = io.read()
print("Provide URL")
local url = io.read()

while true do
  local str = http.get(url.. "/" .. botToken).readAll()
  local obj = json.decode(str)
  for i, v in ipairs(obj.messages) do
    box.sendMessage(v.username.. ": ".. v.text)
  end
  sleep(1)
end
