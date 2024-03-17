package httpRequest

type ContentType string

const (
	AAC    ContentType = "audio/aac"
	ABW    ContentType = "application/x-abiword"
	ARC    ContentType = "application/x-freearc"
	AVI    ContentType = "video/x-msvideo"
	AZW    ContentType = "application/vnd.amazon.ebook"
	BIN    ContentType = "application/octet-stream"
	BZ     ContentType = "application/x-bzip"
	BZ2    ContentType = "application/x-bzip2"
	CSH    ContentType = "application/x-csh"
	CSS    ContentType = "text/css"
	CSV    ContentType = "text/csv"
	DOC    ContentType = "application/msword"
	DOCX   ContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	EOT    ContentType = "application/vnd.ms-fontobject"
	EPUB   ContentType = "application/epub+zip"
	GZ     ContentType = "application/gzip"
	GIF    ContentType = "image/gif"
	HTML   ContentType = "text/html"
	ICO    ContentType = "image/vnd.microsoft.icon"
	ICS    ContentType = "text/calendar"
	JAR    ContentType = "application/java-archive"
	JPEG   ContentType = "image/jpeg"
	JS     ContentType = "text/javascript"
	JSON   ContentType = "application/json"
	JSONLD ContentType = "application/ld+json"
	MID    ContentType = "audio/midi"
	MP3    ContentType = "audio/mpeg"
	MP4    ContentType = "video/mp4"
	MPEG   ContentType = "video/mpeg"
	MPKG   ContentType = "application/vnd.apple.installer+xml"
	ODP    ContentType = "application/vnd.oasis.opendocument.presentation"
	ODS    ContentType = "application/vnd.oasis.opendocument.spreadsheet"
	ODT    ContentType = "application/vnd.oasis.opendocument.text"
	OGA    ContentType = "audio/ogg"
	OGV    ContentType = "video/ogg"
	OGX    ContentType = "application/ogg"
	OTF    ContentType = "font/otf"
	PNG    ContentType = "image/png"
	PDF    ContentType = "application/pdf"
	PPT    ContentType = "application/vnd.ms-powerpoint"
	PPTX   ContentType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	RAR    ContentType = "application/vnd.rar"
	RTF    ContentType = "application/rtf"
	SH     ContentType = "application/x-sh"
	SVG    ContentType = "image/svg+xml"
	SWF    ContentType = "application/x-shockwave-flash"
	TAR    ContentType = "application/x-tar"
	TIFF   ContentType = "image/tiff"
	TTF    ContentType = "font/ttf"
	TXT    ContentType = "text/plain"
	VSD    ContentType = "application/vnd.visio"
	WAV    ContentType = "audio/wav"
	WEBA   ContentType = "audio/webm"
	WEBM   ContentType = "video/webm"
	WEBP   ContentType = "image/webp"
	WOFF   ContentType = "font/woff"
	WOFF2  ContentType = "font/woff2"
	XHTML  ContentType = "application/xhtml+xml"
	XML    ContentType = "application/xml"
	XLS    ContentType = "application/vnd.ms-excel"
	XLSX   ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XUL    ContentType = "application/vnd.mozilla.xul+xml"
	ZIP    ContentType = "application/zip"
	_3GP   ContentType = "video/3gpp"
	_7Z    ContentType = "application/x-7z-compressed"
)

func isValidContentType(content string) bool {
	isValid := map[ContentType]bool{
		AAC:    true,
		ABW:    true,
		ARC:    true,
		AVI:    true,
		AZW:    true,
		BIN:    true,
		BZ:     true,
		BZ2:    true,
		CSH:    true,
		CSS:    true,
		CSV:    true,
		DOC:    true,
		DOCX:   true,
		EOT:    true,
		EPUB:   true,
		GZ:     true,
		GIF:    true,
		HTML:   true,
		ICO:    true,
		ICS:    true,
		JAR:    true,
		JPEG:   true,
		JS:     true,
		JSON:   true,
		JSONLD: true,
		MID:    true,
		MP3:    true,
		MP4:    true,
		MPEG:   true,
		MPKG:   true,
		ODP:    true,
		ODS:    true,
		ODT:    true,
		OGA:    true,
		OGV:    true,
		OGX:    true,
		OTF:    true,
		PNG:    true,
		PDF:    true,
		PPT:    true,
		PPTX:   true,
		RAR:    true,
		RTF:    true,
		SH:     true,
		SVG:    true,
		SWF:    true,
		TAR:    true,
		TIFF:   true,
		TTF:    true,
		TXT:    true,
		VSD:    true,
		WAV:    true,
		WEBA:   true,
		WEBM:   true,
		WEBP:   true,
		WOFF:   true,
		WOFF2:  true,
		XHTML:  true,
		XML:    true,
		XLS:    true,
		XLSX:   true,
		XUL:    true,
		ZIP:    true,
		_3GP:   true,
		_7Z:    true,
	}

	_, ok := isValid[ContentType(content)]
	return ok
}

func (c ContentType) String() string {
	return string(c)
}
