deploy:
	env CGO_ENABLED=0 go build -o proof_master
	scp -P 14922 proof_master admin@36.94.241.4:~/proof_master/
.PHONY: deploy