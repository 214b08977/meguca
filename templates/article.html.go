// Code generated by qtc from "article.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line article.html:1
package templates

//line article.html:1
import "fmt"

//line article.html:2
import "strconv"

//line article.html:3
import "strings"

//line article.html:4
import "github.com/bakape/meguca/common"

//line article.html:5
import "github.com/bakape/meguca/lang"

//line article.html:6
import "github.com/bakape/meguca/imager/assets"

//line article.html:7
import "github.com/bakape/meguca/util"

//line article.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line article.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line article.html:9
func streamrenderArticle(qw422016 *qt422016.Writer, p common.Post, c articleContext) {
//line article.html:10
	id := strconv.FormatUint(p.ID, 10)

//line article.html:11
	ln := lang.Get()

//line article.html:11
	qw422016.N().S(`<article id="p`)
//line article.html:12
	qw422016.N().S(id)
//line article.html:12
	qw422016.N().S(`"`)
//line article.html:12
	qw422016.N().S(` `)
//line article.html:12
	streampostClass(qw422016, p, c.op)
//line article.html:12
	qw422016.N().S(`>`)
//line article.html:13
	streamdeletedToggle(qw422016)
//line article.html:13
	qw422016.N().S(`<header class="spaced"><input type="radio" name="mod-checkbox" class="mod-checkbox hidden">`)
//line article.html:16
	streamrenderSticky(qw422016, c.sticky)
//line article.html:17
	streamrenderLocked(qw422016, c.locked)
//line article.html:18
	if c.subject != "" {
//line article.html:19
		if c.board != "" {
//line article.html:19
			qw422016.N().S(`<b class="board">/`)
//line article.html:21
			qw422016.N().S(c.board)
//line article.html:21
			qw422016.N().S(`/</b>`)
//line article.html:23
		}
//line article.html:23
		qw422016.N().S(`<h3>「`)
//line article.html:25
		qw422016.E().S(c.subject)
//line article.html:25
		qw422016.N().S(`」</h3>`)
//line article.html:27
	}
//line article.html:27
	qw422016.N().S(`<b class="name spaced`)
//line article.html:28
	if p.Auth != common.NotStaff {
//line article.html:28
		qw422016.N().S(` `)
//line article.html:28
		qw422016.N().S(`admin`)
//line article.html:28
	}
//line article.html:28
	if p.Sage {
//line article.html:28
		qw422016.N().S(` `)
//line article.html:28
		qw422016.N().S(`sage`)
//line article.html:28
	}
//line article.html:28
	qw422016.N().S(`">`)
//line article.html:29
	if p.Name != "" || p.Trip == "" {
//line article.html:29
		qw422016.N().S(`<span>`)
//line article.html:31
		if p.Name != "" {
//line article.html:32
			qw422016.E().S(p.Name)
//line article.html:33
		} else {
//line article.html:34
			qw422016.N().S(ln.Common.Posts["anon"])
//line article.html:35
		}
//line article.html:35
		qw422016.N().S(`</span>`)
//line article.html:37
	}
//line article.html:38
	if p.Trip != "" {
//line article.html:38
		qw422016.N().S(`<code>!`)
//line article.html:40
		qw422016.E().S(p.Trip)
//line article.html:40
		qw422016.N().S(`</code>`)
//line article.html:42
	}
//line article.html:43
	if p.Auth != common.NotStaff {
//line article.html:43
		qw422016.N().S(`<span>##`)
//line article.html:45
		qw422016.N().S(` `)
//line article.html:45
		qw422016.N().S(ln.Common.Posts[p.Auth.String()])
//line article.html:45
		qw422016.N().S(`</span>`)
//line article.html:47
	}
//line article.html:47
	qw422016.N().S(`</b>`)
//line article.html:49
	if p.Flag != "" {
//line article.html:50
		title, ok := countryMap[p.Flag]

//line article.html:51
		if !ok {
//line article.html:52
			title = p.Flag

//line article.html:53
		}
//line article.html:54
		if strings.HasPrefix(p.Flag, "us-") {
//line article.html:55
			title2, ok2 := countryMap["us"]

//line article.html:56
			if !ok2 {
//line article.html:57
				title2 = "us"

//line article.html:58
			}
//line article.html:58
			qw422016.N().S(`<img class="flag" src="/assets/flags/us.svg" title="`)
//line article.html:59
			qw422016.N().S(title2)
//line article.html:59
			qw422016.N().S(`">`)
//line article.html:60
		}
//line article.html:60
		qw422016.N().S(`<img class="flag" src="/assets/flags/`)
//line article.html:61
		qw422016.N().S(p.Flag)
//line article.html:61
		qw422016.N().S(`.svg" title="`)
//line article.html:61
		qw422016.N().S(title)
//line article.html:61
		qw422016.N().S(`">`)
//line article.html:62
	}
//line article.html:62
	qw422016.N().S(`<time>`)
//line article.html:64
	qw422016.N().S(formatTime(p.Time))
//line article.html:64
	qw422016.N().S(`</time><nav>`)
//line article.html:67
	url := "#p" + id

//line article.html:68
	if c.index {
//line article.html:69
		url = util.ConcatStrings("/all/", id, "?last=100", url)

//line article.html:70
	}
//line article.html:70
	qw422016.N().S(`<a href="`)
//line article.html:71
	qw422016.N().S(url)
//line article.html:71
	qw422016.N().S(`">No.</a><a class="quote">`)
//line article.html:75
	qw422016.N().S(id)
//line article.html:75
	qw422016.N().S(`</a></nav>`)
//line article.html:78
	if c.index && c.subject != "" {
//line article.html:78
		qw422016.N().S(`<span>`)
//line article.html:80
		streamexpandLink(qw422016, "all", id)
//line article.html:81
		streamlast100Link(qw422016, "all", id)
//line article.html:81
		qw422016.N().S(`</span>`)
//line article.html:83
	}
//line article.html:84
	streamcontrolLink(qw422016)
//line article.html:85
	if c.op == p.ID {
//line article.html:86
		streamthreadWatcherToggle(qw422016, p.ID)
//line article.html:87
	}
//line article.html:87
	qw422016.N().S(`</header>`)
//line article.html:89
	var src string

//line article.html:90
	if p.Image != nil {
//line article.html:91
		img := *p.Image

//line article.html:92
		src = assets.SourcePath(img.FileType, img.SHA1)

//line article.html:92
		qw422016.N().S(`<figcaption class="spaced"><a class="image-toggle act" hidden></a><span class="spaced image-search-container">`)
//line article.html:96
		streamimageSearch(qw422016, c.root, img)
//line article.html:96
		qw422016.N().S(`</span><span class="fileinfo">`)
//line article.html:99
		if img.Audio {
//line article.html:99
			qw422016.N().S(`<span>♫</span>`)
//line article.html:103
		}
//line article.html:104
		if img.Length != 0 {
//line article.html:104
			qw422016.N().S(`<span>`)
//line article.html:106
			l := img.Length

//line article.html:107
			if l < 60 {
//line article.html:108
				qw422016.N().S(fmt.Sprintf("0:%02d", l))
//line article.html:109
			} else {
//line article.html:110
				min := l / 60

//line article.html:111
				qw422016.N().S(fmt.Sprintf("%02d:%02d", min, l-min*60))
//line article.html:112
			}
//line article.html:112
			qw422016.N().S(`</span>`)
//line article.html:114
		}
//line article.html:114
		qw422016.N().S(`<span>`)
//line article.html:116
		qw422016.N().S(readableFileSize(img.Size))
//line article.html:116
		qw422016.N().S(`</span>`)
//line article.html:118
		if img.Dims != [4]uint16{} {
//line article.html:118
			qw422016.N().S(`<span>`)
//line article.html:120
			qw422016.N().S(strconv.FormatUint(uint64(img.Dims[0]), 10))
//line article.html:120
			qw422016.N().S(`x`)
//line article.html:122
			qw422016.N().S(strconv.FormatUint(uint64(img.Dims[1]), 10))
//line article.html:122
			qw422016.N().S(`</span>`)
//line article.html:124
		}
//line article.html:125
		if img.Artist != "" {
//line article.html:125
			qw422016.N().S(`<span>`)
//line article.html:127
			qw422016.E().S(img.Artist)
//line article.html:127
			qw422016.N().S(`</span>`)
//line article.html:129
		}
//line article.html:130
		if img.Title != "" {
//line article.html:130
			qw422016.N().S(`<span>`)
//line article.html:132
			qw422016.E().S(img.Title)
//line article.html:132
			qw422016.N().S(`</span>`)
//line article.html:134
		}
//line article.html:134
		qw422016.N().S(`</span>`)
//line article.html:136
		name := imageName(img.FileType, img.Name)

//line article.html:136
		qw422016.N().S(`<a href="`)
//line article.html:137
		qw422016.N().S(assets.RelativeSourcePath(img.FileType, img.SHA1))
//line article.html:137
		qw422016.N().S(`" download="`)
//line article.html:137
		qw422016.N().S(name)
//line article.html:137
		qw422016.N().S(`">`)
//line article.html:138
		qw422016.N().S(name)
//line article.html:138
		qw422016.N().S(`</a></figcaption>`)
//line article.html:141
	}
//line article.html:141
	qw422016.N().S(`<div class="post-container">`)
//line article.html:143
	if p.Image != nil {
//line article.html:144
		img := *p.Image

//line article.html:144
		qw422016.N().S(`<figure><a target="_blank" href="`)
//line article.html:146
		qw422016.N().S(src)
//line article.html:146
		qw422016.N().S(`">`)
//line article.html:147
		switch {
//line article.html:148
		case img.ThumbType == common.NoFile:
//line article.html:149
			var file string

//line article.html:150
			switch img.FileType {
//line article.html:151
			case common.WEBM, common.MP4, common.MP3, common.OGG, common.FLAC:
//line article.html:152
				file = "audio"

//line article.html:153
			default:
//line article.html:154
				file = "file"

//line article.html:155
			}
//line article.html:155
			qw422016.N().S(`<img src="/assets/`)
//line article.html:156
			qw422016.N().S(file)
//line article.html:156
			qw422016.N().S(`.png" width="150" height="150" loading="lazy">`)
//line article.html:157
		case img.Spoiler:
//line article.html:160
			qw422016.N().S(`<img src="/assets/spoil/default.jpg" width="150" height="150" loading="lazy">`)
//line article.html:162
		default:
//line article.html:162
			qw422016.N().S(`<img src="`)
//line article.html:163
			qw422016.N().S(assets.ThumbPath(img.ThumbType, img.SHA1))
//line article.html:163
			qw422016.N().S(`" width="`)
//line article.html:163
			qw422016.N().D(int(img.Dims[2]))
//line article.html:163
			qw422016.N().S(`" height="`)
//line article.html:163
			qw422016.N().D(int(img.Dims[3]))
//line article.html:163
			qw422016.N().S(`" loading="lazy">`)
//line article.html:164
		}
//line article.html:164
		qw422016.N().S(`</a></figure>`)
//line article.html:167
	}
//line article.html:167
	qw422016.N().S(`<blockquote>`)
//line article.html:169
	streambody(qw422016, p, c.op, c.board, c.index, c.rbText, c.pyu)
//line article.html:169
	qw422016.N().S(`</blockquote>`)
//line article.html:171
	for _, e := range p.Moderation {
//line article.html:171
		qw422016.N().S(`<b class="admin post-moderation">`)
//line article.html:173
		streampostModeration(qw422016, e)
//line article.html:173
		qw422016.N().S(`<br></b>`)
//line article.html:176
	}
//line article.html:176
	qw422016.N().S(`</div>`)
//line article.html:178
	if c.omit != 0 {
//line article.html:178
		qw422016.N().S(`<span class="omit spaced" data-omit="`)
//line article.html:179
		qw422016.N().D(c.omit)
//line article.html:179
		qw422016.N().S(`" data-image-omit="`)
//line article.html:179
		qw422016.N().D(c.imageOmit)
//line article.html:179
		qw422016.N().S(`">`)
//line article.html:180
		if c.imageOmit == 0 {
//line article.html:181
			qw422016.N().S(fmt.Sprintf(ln.Common.Format["postsOmitted"], c.omit))
//line article.html:182
		} else {
//line article.html:183
			qw422016.N().S(fmt.Sprintf(ln.Common.Format["postsAndImagesOmitted"], c.omit, c.imageOmit))
//line article.html:184
		}
//line article.html:184
		qw422016.N().S(`<span class="act"><a href="`)
//line article.html:186
		qw422016.N().S(strconv.FormatUint(c.op, 10))
//line article.html:186
		qw422016.N().S(`">`)
//line article.html:187
		qw422016.N().S(ln.Common.Posts["seeAll"])
//line article.html:187
		qw422016.N().S(`</a></span></span>`)
//line article.html:191
	}
//line article.html:192
	if bls := c.backlinks[p.ID]; len(bls) != 0 {
//line article.html:192
		qw422016.N().S(`<span class="backlinks spaced">`)
//line article.html:194
		for _, l := range bls {
//line article.html:194
			qw422016.N().S(`<em>`)
//line article.html:196
			streampostLink(qw422016, l, c.index || l.OP != c.op, c.index)
//line article.html:196
			qw422016.N().S(`</em>`)
//line article.html:198
		}
//line article.html:198
		qw422016.N().S(`</span>`)
//line article.html:200
	}
//line article.html:200
	qw422016.N().S(`</article>`)
//line article.html:202
}

//line article.html:202
func writerenderArticle(qq422016 qtio422016.Writer, p common.Post, c articleContext) {
//line article.html:202
	qw422016 := qt422016.AcquireWriter(qq422016)
//line article.html:202
	streamrenderArticle(qw422016, p, c)
//line article.html:202
	qt422016.ReleaseWriter(qw422016)
//line article.html:202
}

//line article.html:202
func renderArticle(p common.Post, c articleContext) string {
//line article.html:202
	qb422016 := qt422016.AcquireByteBuffer()
//line article.html:202
	writerenderArticle(qb422016, p, c)
//line article.html:202
	qs422016 := string(qb422016.B)
//line article.html:202
	qt422016.ReleaseByteBuffer(qb422016)
//line article.html:202
	return qs422016
//line article.html:202
}

// Render image search links according to file type

//line article.html:205
func streamimageSearch(qw422016 *qt422016.Writer, root string, img common.Image) {
//line article.html:206
	if img.ThumbType == common.NoFile || img.FileType == common.PDF {
//line article.html:207
		return
//line article.html:208
	}
//line article.html:210
	url := root + assets.ImageSearchPath(img.ImageCommon)

//line article.html:210
	qw422016.N().S(`<a class="image-search google" target="_blank" rel="nofollow" href="https://www.google.com/searchbyimage?image_url=`)
//line article.html:211
	qw422016.N().S(url)
//line article.html:211
	qw422016.N().S(`">G</a><a class="image-search yandex" target="_blank" rel="nofollow" href="https://yandex.com/images/search?source=collections&rpt=imageview&url=`)
//line article.html:214
	qw422016.N().S(url)
//line article.html:214
	qw422016.N().S(`">Yd</a><a class="image-search iqdb" target="_blank" rel="nofollow" href="http://iqdb.org/?url=`)
//line article.html:217
	qw422016.N().S(url)
//line article.html:217
	qw422016.N().S(`">Iq</a><a class="image-search saucenao" target="_blank" rel="nofollow" href="http://saucenao.com/search.php?db=999&url=`)
//line article.html:220
	qw422016.N().S(url)
//line article.html:220
	qw422016.N().S(`">Sn</a><a class="image-search tracemoe" target="_blank" rel="nofollow" href="https://trace.moe/?url=`)
//line article.html:223
	qw422016.N().S(url)
//line article.html:223
	qw422016.N().S(`">Tm</a>`)
//line article.html:226
	switch img.FileType {
//line article.html:227
	case common.JPEG, common.PNG, common.GIF, common.WEBM:
//line article.html:227
		qw422016.N().S(`<a class="image-search desuarchive" target="_blank" rel="nofollow" href="https://desuarchive.org/_/search/image/`)
//line article.html:228
		qw422016.N().S(img.MD5)
//line article.html:228
		qw422016.N().S(`">Da</a>`)
//line article.html:231
	}
//line article.html:232
	switch img.FileType {
//line article.html:233
	case common.JPEG, common.PNG:
//line article.html:233
		qw422016.N().S(`<a class="image-search exhentai" target="_blank" rel="nofollow" href="http://exhentai.org/?fs_similar=1&fs_exp=1&f_shash=`)
//line article.html:234
		qw422016.N().S(img.SHA1)
//line article.html:234
		qw422016.N().S(`">Ex</a>`)
//line article.html:237
	}
//line article.html:238
}

//line article.html:238
func writeimageSearch(qq422016 qtio422016.Writer, root string, img common.Image) {
//line article.html:238
	qw422016 := qt422016.AcquireWriter(qq422016)
//line article.html:238
	streamimageSearch(qw422016, root, img)
//line article.html:238
	qt422016.ReleaseWriter(qw422016)
//line article.html:238
}

//line article.html:238
func imageSearch(root string, img common.Image) string {
//line article.html:238
	qb422016 := qt422016.AcquireByteBuffer()
//line article.html:238
	writeimageSearch(qb422016, root, img)
//line article.html:238
	qs422016 := string(qb422016.B)
//line article.html:238
	qt422016.ReleaseByteBuffer(qb422016)
//line article.html:238
	return qs422016
//line article.html:238
}
