local printer =     require("printer")
local http    =     require("http")
local url     =     require("url")
local net     =     require("net")
local client  =     http.client()

script = {
    title = "Honey pot checker",
    author = "krishpranav (wprecon)",
    risklevel = "low",
    type = "Checker",
    description = "It will check it the target is a honey pot, and give a percentage based on shodan"
    references = {""},
}

function main(target)
    local uri_host = url.host(target)
    local ip = net.lookup_ip(uri_host)

    local request = http.request("GET", "https://api.shodan.io/labs/honeyscore/"..ip.."?key=C23OXE0bVMrul2YeqcL7zxb6jZ4pj2by")
    local response, err = client:do_request(request)

    if err then
        printer.danger(err)
    end

    if response_code == 200 then
        printer.done("With a "..convert(response.body).." chance of this host being a Honeypot.")
    end
end


function convert(text)
    text = text:gsub("0.0", "0%")
    text = text:gsub("0.1", "10%")
    text = text:gsub("0.2", "20%")
    text = text:gsub("0.3", "30%")
    text = text:gsub("0.4", "40%")
    text = text:gsub("0.5", "50%")
    text = text:gsub("0.6", "60%")
    text = text:gsub("0.7", "70%")
    text = text:gsub("0.8", "80%")
    text = text:gsub("0.9", "90%")
    text = text:gsub("1.0", "100%")

    return text
end