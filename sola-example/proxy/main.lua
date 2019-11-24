function handle()
	if (URL == "/403")
	then
		return 403, "403"
    else
        index = string.find(URL, "/err")
        if (index == 1)
        then
            e = string.gsub(URL, "/err", "", 1)
            error("error: "..e)
        end
	end
end
