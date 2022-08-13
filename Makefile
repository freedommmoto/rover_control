test:
	go test -v -cover ./...
newfile:
	rm -f ./instructions_file/instructions.txt
	echo "24\nR\nF\nL\nF\nL\nL\nF\nR" >> ./instructions_file/instructions.txt