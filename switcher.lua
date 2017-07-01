keyboardWatcher = hs.usb.watcher.new(function(params)
    if params["productName"] == "Realforce 87" and params["eventType"] == "added" then
      hs.notify.new({title="Karabiner-Elements Profile Switcher", informativeText=params["productName"].." connected"}):send()
      hs.task.new("/Users/uraura/bin/karabiner-elements-profile-switcher", nil, {"use", "RealForce"}):start()
    elseif params["productName"] == "Realforce 87" and params["eventType"] == "removed" then
      hs.notify.new({title="Karabiner-Elements Profile Switcher", informativeText=params["productName"].." disconnected"}):send()
      hs.task.new("/Users/uraura/bin/karabiner-elements-profile-switcher", nil, {"use", "Default"}):start()
    end
end)
keyboardWatcher:start()
