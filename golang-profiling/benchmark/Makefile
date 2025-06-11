pprof-ui:
	curl -v -H "Authorization: foobar" -o profile.pb.gz \
	http://localhost:8080/debug/pprof/profile\?seconds\=60
	go tool pprof -http=:8099 profile.pb.gz


# curl -o cpu.pb.gz http://localhost:8080/debug/pprof/profile?seconds=30

# to visualize
# go tool pprof cpu.pb.gz
# (pprof) top
# (pprof) web


# for web ui
go tool pprof -http=:7070 cpu.pb.gz


# go tool pprof -http=:8099 profile.pb.gz

