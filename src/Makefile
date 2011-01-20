
all: searcher 

SRC= main.go docmap.go querrymap.go utils.go invert.go stemmer.go qproc.go webserver.go
LN= main.8


GC = 8g
GL = 8l

$(LN): $(SRC)
	$(GC) -o $@ $(SRC)

searcher: $(LN)
	$(GL) -o $@ $(LN)

