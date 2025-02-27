local inspect = require "inspect"

local function print_headers(request_headers)
  for key, value in pairs(request_headers) do
    print(string.format(" - %s: %s", tostring(key), tostring(value)))
  end
end

function on_envoy_request(request_handle)
  local request_headers = request_handle:headers()
  print_headers(request_headers)
end

function envoy_on_response(response_handle)
  response_handle:
end