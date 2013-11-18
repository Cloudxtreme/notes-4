**us klavyede turkce karakterleri kullanabilmek icin**

```bash
setxkbmap tr -variant alt
```

**south(django) icin alias**

```bash
function south_first() {
	./manage.py schemamigration --initial "$@";
	./manage.py migrate "$@";
}

function south_update() {
	./manage.py schemamigration --auto "$@";
	./manage.py migrate "$@";
}
```

