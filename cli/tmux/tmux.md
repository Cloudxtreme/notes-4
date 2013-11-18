## başlangıç

**yeni bir oturum başlatmak**

    tmux new -s oturum_adi

**açılan oturumu "detach" etmek.**

(görünürde çıkıyoruz ama arkaplanda komut(lar) çalışmaya devam ediyor.)

    [PREFIX] d
(PREFIX standart olarak CTRL+B, ben CTRL+A olarak değiştirdim. notların tamamında [PREFIX] olarak referans edilecektir.)

**arkaplandaki oturuma devam etmek**

    tmux attach -t oturum_adi (tek bir oturum varsa tmux attach)

**örnek:**

    tmux new -s genel
    top
    [PREFIX] d
    tmux list-sessions # açık oturumları listeler, "tmux ls" de kullanılabilir.
    tmux attach -t genel

**oturumu öldürmek**

takılan (hang olmuş) bir oturum varsa:

    tmux kill-session -t target_name

## pencereler

pencereler, tab olarak düşünülebilir.

**tmux içindeyken yeni bir pencere yaratmak için:**

    [PREFIX] + c # c for create

**bulunulan pencerenin adını değiştirmek**

    [PREFIX] + : rename-window

**pencereler arasında gezinmek**

    [PREFIX] + n

veya

    [PREFIX] + p

n(ext) sonraki, (p)revious önceki pencereye geçiş sağlar.

veya

    [PREFIX] + PENCERE_NUMARASI

0 ile 9 arasında olabilir.


**pencereyi kapatmak**

    exit

##panellerle çalışmak (panes, dilimler?)

pencereleri X kadar panele bölmek mümkün.

**pencereyi "horizantal" bölmek**

    [PREFIX] + %

**pencereyi "vertical" bölmek**

    [PREFIX] + "

~/.tmux.conf'a eklenecek şu satırlarla daha "semantik" bir şekilde halloluyor.

    bind | split-window -h
    bind - split-window -v

**paneller arası gezinti**

    [PREFIX] + o

veya

    [PREFIX] + [yön tuşları]

**hazır panel planları**

tmux ile beraber geliyor:

* horizontal
* even-vertical
* main-horizontal
* main
* tiled

**bunlar arasında geçiş yapmak için:**

    [PREFIX] + [space]

**panel numaralarını görmek (bir saniyeliğine)**

    [PREFIX] + q

**panelden çıkış**

    exit

## tmux komut modu

    [PREFIX] + :

ile komut moduna giriliyor. yazılabilecek komutlar çok fazla. list-commands komutuyla ilgili komutlar geliyor.

## tmux ayarlarıyla oynamak

**prefix'i değiştirmek**

    touch ~/.tmux.conf

ile kullanıcınıza ait ayar dosyasını oluşturun.

prefix her komutta var, CTRL+B yerine CTRL+a daha rahat gibi.

    set -g prefix C-a
    unbind C-b

tmux.conf'ta yapılan değişikliklerden sonra değişiklerin yansıması için komut modunda:

    source-file ~/.tmux.conf

####**pencere ve panellerde index'i 1'den başlatmak**

bilgisayarlar 0'dan başlayabilir. biz 1'den başlasak daha iyi.

    set -g base-index 1
    setw -g pane-base-index 1

**source-file komutu için kısayol**

    bind r source-file ~/.tmux.conf \; display-message "Config reloaded..."

**mouse desteğini aktif hale getirmek**

(kısayollara alışırsanız daha iyi/hızlı olur uzun vadede.)

    setw -g mode-mouse on
    set -g mouse-select-pane on
    set -g mouse-resize-pane on
    set -g mouse-select-window on



`