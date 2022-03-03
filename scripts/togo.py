with open("words.txt") as f:
    with open("words.go", "w") as fw:
        fw.write("package main\n\n")
        fw.write("// SOURCE: https://github.com/dwyl/english-words\n\n")
        fw.write("var WORDS = []string{\n")

        for line in f:
            fw.write(f'"{line.strip()}",\n')

        fw.write("}")
