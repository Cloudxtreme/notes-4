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

**poor man's backup solution*

```bash
mysqldump -e -u user -pparolan dbname | gzip | uuencode dbbackup_e.gz | mail target_email
```
