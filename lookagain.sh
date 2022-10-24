find . -name "*.sh" | sort -r | rev | cut -d "/" -f1 | rev | sed 's/.sh//g'
