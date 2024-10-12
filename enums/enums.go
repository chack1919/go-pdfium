package enums

// The file identifier entry type. See section 14.4 "File Identifiers" of the
// ISO 32000-1:2008 spec.
type FPDF_FILEIDTYPE int

const (
	FPDF_FILEIDTYPE_PERMANENT FPDF_FILEIDTYPE = 0
	FPDF_FILEIDTYPE_CHANGING  FPDF_FILEIDTYPE = 1
)

type FPDF_ACTION_ACTION uint32

const (
	FPDF_ACTION_ACTION_UNSUPPORTED  FPDF_ACTION_ACTION = 0 // Action type: unsupported action type.
	FPDF_ACTION_ACTION_GOTO         FPDF_ACTION_ACTION = 1 // This action contains information which can be used to go to a destination within current document.
	FPDF_ACTION_ACTION_REMOTEGOTO   FPDF_ACTION_ACTION = 2 // This action contains information which can be used to launch an application or opens or prints a document.
	FPDF_ACTION_ACTION_URI          FPDF_ACTION_ACTION = 3 // This action contains information which can be used to go to a destination within another document.
	FPDF_ACTION_ACTION_LAUNCH       FPDF_ACTION_ACTION = 4 // This action contains information which identifies (resolves to) a resource on the Internet - such as web pages, a file that is the destination of a hypertext link, and etc.
	FPDF_ACTION_ACTION_EMBEDDEDGOTO FPDF_ACTION_ACTION = 5 // This action contains information which can be used to Go to a destination in an embedded file.
)

type FPDF_PAGE_ROTATION int

const (
	FPDF_PAGE_ROTATION_NONE   FPDF_PAGE_ROTATION = 0 // 0: no rotation.
	FPDF_PAGE_ROTATION_90_CW  FPDF_PAGE_ROTATION = 1 // 1: rotate 90 degrees in clockwise direction.
	FPDF_PAGE_ROTATION_180_CW FPDF_PAGE_ROTATION = 2 // 2: rotate 180 degrees in clockwise direction.
	FPDF_PAGE_ROTATION_270_CW FPDF_PAGE_ROTATION = 3 // 3: rotate 270 degrees in clockwise direction.
)

// View destination fit types. See pdfmark reference v9, page 48.
type FPDF_PDFDEST_VIEW uint32

const (
	FPDF_PDFDEST_VIEW_UNKNOWN_MODE FPDF_PDFDEST_VIEW = 0
	FPDF_PDFDEST_VIEW_XYZ          FPDF_PDFDEST_VIEW = 1
	FPDF_PDFDEST_VIEW_FIT          FPDF_PDFDEST_VIEW = 2
	FPDF_PDFDEST_VIEW_FITH         FPDF_PDFDEST_VIEW = 3
	FPDF_PDFDEST_VIEW_FITV         FPDF_PDFDEST_VIEW = 4
	FPDF_PDFDEST_VIEW_FITR         FPDF_PDFDEST_VIEW = 5
	FPDF_PDFDEST_VIEW_FITB         FPDF_PDFDEST_VIEW = 6
	FPDF_PDFDEST_VIEW_FITBH        FPDF_PDFDEST_VIEW = 7
	FPDF_PDFDEST_VIEW_FITBV        FPDF_PDFDEST_VIEW = 8
)

// Additional-action types of page object
type FPDF_PAGE_AACTION int

const (
	FPDF_PAGE_AACTION_OPEN  FPDF_PAGE_AACTION = 0 // OPEN (/O) -- An action to be performed when the page is opened
	FPDF_PAGE_AACTION_CLOSE FPDF_PAGE_AACTION = 1 // CLOSE (/C) -- An action to be performed when the page is closed
)

// Additional actions type of document
type FPDF_DOC_AACTION int

const (
	FPDF_DOC_AACTION_WC FPDF_DOC_AACTION = 0x10 // WC, before closing document, JavaScript action.
	FPDF_DOC_AACTION_WS FPDF_DOC_AACTION = 0x11 // WS before saving document, JavaScript action.
	FPDF_DOC_AACTION_DS FPDF_DOC_AACTION = 0x12 // DS, after saving document, JavaScript action.
	FPDF_DOC_AACTION_WP FPDF_DOC_AACTION = 0x13 // WP, before printing document, JavaScript action.
	FPDF_DOC_AACTION_DP FPDF_DOC_AACTION = 0x14 // DP, after printing document, JavaScript action.
)

type FPDF_UNSP int

const (
	FPDF_UNSP_DOC_XFAFORM               FPDF_UNSP = 1
	FPDF_UNSP_DOC_PORTABLECOLLECTION    FPDF_UNSP = 2
	FPDF_UNSP_DOC_ATTACHMENT            FPDF_UNSP = 3
	FPDF_UNSP_DOC_SECURITY              FPDF_UNSP = 4
	FPDF_UNSP_DOC_SHAREDREVIEW          FPDF_UNSP = 5
	FPDF_UNSP_DOC_SHAREDFORM_ACROBAT    FPDF_UNSP = 6
	FPDF_UNSP_DOC_SHAREDFORM_FILESYSTEM FPDF_UNSP = 7
	FPDF_UNSP_DOC_SHAREDFORM_EMAIL      FPDF_UNSP = 8
	FPDF_UNSP_ANNOT_3DANNOT             FPDF_UNSP = 11
	FPDF_UNSP_ANNOT_MOVIE               FPDF_UNSP = 12
	FPDF_UNSP_ANNOT_SOUND               FPDF_UNSP = 13
	FPDF_UNSP_ANNOT_SCREEN_MEDIA        FPDF_UNSP = 14
	FPDF_UNSP_ANNOT_SCREEN_RICHMEDIA    FPDF_UNSP = 15
	FPDF_UNSP_ANNOT_ATTACHMENT          FPDF_UNSP = 16
	FPDF_UNSP_ANNOT_SIG                 FPDF_UNSP = 17
)

type FPDF_FXFONT_CHARSET int

const (
	FPDF_FXFONT_ANSI_CHARSET            FPDF_FXFONT_CHARSET = 0
	FPDF_FXFONT_DEFAULT_CHARSET         FPDF_FXFONT_CHARSET = 1
	FPDF_FXFONT_SYMBOL_CHARSET          FPDF_FXFONT_CHARSET = 2
	FPDF_FXFONT_SHIFTJIS_CHARSET        FPDF_FXFONT_CHARSET = 128
	FPDF_FXFONT_HANGEUL_CHARSET         FPDF_FXFONT_CHARSET = 129
	FPDF_FXFONT_GB2312_CHARSET          FPDF_FXFONT_CHARSET = 134
	FPDF_FXFONT_CHINESEBIG5_CHARSET     FPDF_FXFONT_CHARSET = 136
	FPDF_FXFONT_GREEK_CHARSET           FPDF_FXFONT_CHARSET = 161
	FPDF_FXFONT_VIETNAMESE_CHARSET      FPDF_FXFONT_CHARSET = 163
	FPDF_FXFONT_HEBREW_CHARSET          FPDF_FXFONT_CHARSET = 177
	FPDF_FXFONT_ARABIC_CHARSET          FPDF_FXFONT_CHARSET = 178
	FPDF_FXFONT_CYRILLIC_CHARSET        FPDF_FXFONT_CHARSET = 204
	FPDF_FXFONT_THAI_CHARSET            FPDF_FXFONT_CHARSET = 22
	FPDF_FXFONT_EASTERNEUROPEAN_CHARSET FPDF_FXFONT_CHARSET = 238
)

type FPDF_OBJECT_TYPE int

const (
	FPDF_OBJECT_TYPE_UNKNOWN    FPDF_OBJECT_TYPE = 0
	FPDF_OBJECT_TYPE_BOOLEAN    FPDF_OBJECT_TYPE = 1
	FPDF_OBJECT_TYPE_NUMBER     FPDF_OBJECT_TYPE = 2
	FPDF_OBJECT_TYPE_STRING     FPDF_OBJECT_TYPE = 3
	FPDF_OBJECT_TYPE_NAME       FPDF_OBJECT_TYPE = 4
	FPDF_OBJECT_TYPE_ARRAY      FPDF_OBJECT_TYPE = 5
	FPDF_OBJECT_TYPE_DICTIONARY FPDF_OBJECT_TYPE = 6
	FPDF_OBJECT_TYPE_STREAM     FPDF_OBJECT_TYPE = 7
	FPDF_OBJECT_TYPE_NULLOBJ    FPDF_OBJECT_TYPE = 8
	FPDF_OBJECT_TYPE_REFERENCE  FPDF_OBJECT_TYPE = 9
)

type FPDF_TEXT_RENDERMODE int

const (
	FPDF_TEXTRENDERMODE_UNKNOWN          FPDF_TEXT_RENDERMODE = -1
	FPDF_TEXTRENDERMODE_FILL             FPDF_TEXT_RENDERMODE = 0
	FPDF_TEXTRENDERMODE_STROKE           FPDF_TEXT_RENDERMODE = 1
	FPDF_TEXTRENDERMODE_FILL_STROKE      FPDF_TEXT_RENDERMODE = 2
	FPDF_TEXTRENDERMODE_INVISIBLE        FPDF_TEXT_RENDERMODE = 3
	FPDF_TEXTRENDERMODE_FILL_CLIP        FPDF_TEXT_RENDERMODE = 4
	FPDF_TEXTRENDERMODE_STROKE_CLIP      FPDF_TEXT_RENDERMODE = 5
	FPDF_TEXTRENDERMODE_FILL_STROKE_CLIP FPDF_TEXT_RENDERMODE = 6
	FPDF_TEXTRENDERMODE_CLIP             FPDF_TEXT_RENDERMODE = 7
)

type FPDF_BITMAP_FORMAT int

const (
	FPDF_BITMAP_FORMAT_UNKNOWN FPDF_BITMAP_FORMAT = 0 // Unknown or unsupported format.
	FPDF_BITMAP_FORMAT_GRAY    FPDF_BITMAP_FORMAT = 1 // Gray scale bitmap, one byte per pixel.
	FPDF_BITMAP_FORMAT_BGR     FPDF_BITMAP_FORMAT = 2 // 3 bytes per pixel, byte order: blue, green, red.
	FPDF_BITMAP_FORMAT_BGRX    FPDF_BITMAP_FORMAT = 3 // 4 bytes per pixel, byte order: blue, green, red, unused.
	FPDF_BITMAP_FORMAT_BGRA    FPDF_BITMAP_FORMAT = 4 // 4 bytes per pixel, byte order: blue, green, red, alpha.
)

type FPDF_DUPLEXTYPE int

const (
	FPDF_DUPLEXTYPE_UNDEFINED              FPDF_DUPLEXTYPE = 0
	FPDF_DUPLEXTYPE_SIMPLEX                FPDF_DUPLEXTYPE = 1
	FPDF_DUPLEXTYPE_DUPLEX_FLIP_SHORT_EDGE FPDF_DUPLEXTYPE = 2
	FPDF_DUPLEXTYPE_DUPLEX_FLIP_LONG_EDGE  FPDF_DUPLEXTYPE = 3
)

type FPDF_RENDER_FLAG int

const (
	FPDF_RENDER_FLAG_ANNOT                    FPDF_RENDER_FLAG = 0x01   // Set if annotations are to be rendered.
	FPDF_RENDER_FLAG_LCD_TEXT                 FPDF_RENDER_FLAG = 0x02   // Set if using text rendering optimized for LCD display. This flag will only take effect if anti-aliasing is enabled for text.
	FPDF_RENDER_FLAG_NO_NATIVETEXT            FPDF_RENDER_FLAG = 0x04   // Don't use the native text output available on some platforms.
	FPDF_RENDER_FLAG_GRAYSCALE                FPDF_RENDER_FLAG = 0x08   // Grayscale output.
	FPDF_RENDER_FLAG_DEBUG_INFO               FPDF_RENDER_FLAG = 0x80   // Obsolete, has no effect, retained for compatibility.
	FPDF_RENDER_FLAG_NO_CATCH                 FPDF_RENDER_FLAG = 0x100  // Obsolete, has no effect, retained for compatibility.
	FPDF_RENDER_FLAG_RENDER_LIMITEDIMAGECACHE FPDF_RENDER_FLAG = 0x200  // Limit image cache size.
	FPDF_RENDER_FLAG_RENDER_FORCEHALFTONE     FPDF_RENDER_FLAG = 0x400  // Always use halftone for image stretching.
	FPDF_RENDER_FLAG_PRINTING                 FPDF_RENDER_FLAG = 0x800  // Render for printing.
	FPDF_RENDER_FLAG_RENDER_NO_SMOOTHTEXT     FPDF_RENDER_FLAG = 0x1000 // Set to disable anti-aliasing on text. This flag will also disable LCD optimization for text rendering.
	FPDF_RENDER_FLAG_RENDER_NO_SMOOTHIMAGE    FPDF_RENDER_FLAG = 0x2000 // Set to disable anti-aliasing on images.
	FPDF_RENDER_FLAG_RENDER_NO_SMOOTHPATH     FPDF_RENDER_FLAG = 0x4000 // Set to disable anti-aliasing on paths.
	FPDF_RENDER_FLAG_REVERSE_BYTE_ORDER       FPDF_RENDER_FLAG = 0x10   // Set whether to render in a reverse Byte order, this flag is only used when rendering to a bitmap.
	FPDF_RENDER_FLAG_CONVERT_FILL_TO_STROKE   FPDF_RENDER_FLAG = 0x20   // Set whether fill paths need to be stroked. This flag is only used when FPDF_COLORSCHEME is passed in, since with a single fill color for paths the boundaries of adjacent fill paths are less visible.
)

type FPDF_PRINTMODE int

const (
	FPDF_PRINTMODE_EMF                            FPDF_PRINTMODE = 0 // To output EMF (default)
	FPDF_PRINTMODE_TEXTONLY                       FPDF_PRINTMODE = 1 // to output text only (for charstream devices)
	FPDF_PRINTMODE_POSTSCRIPT2                    FPDF_PRINTMODE = 2 // to output level 2 PostScript into EMF as a series of GDI comments.
	FPDF_PRINTMODE_POSTSCRIPT3                    FPDF_PRINTMODE = 3 // to output level 3 PostScript into EMF as a series of GDI comments.
	FPDF_PRINTMODE_POSTSCRIPT2_PASSTHROUGH        FPDF_PRINTMODE = 4 // to output level 2 PostScript via ExtEscape() in PASSTHROUGH mode.
	FPDF_PRINTMODE_POSTSCRIPT3_PASSTHROUGH        FPDF_PRINTMODE = 5 // to output level 3 PostScript via ExtEscape() in PASSTHROUGH mode.
	FPDF_PRINTMODE_EMF_IMAGE_MASKS                FPDF_PRINTMODE = 6 // to output EMF, with more efficient processing of documents containing image masks.
	FPDF_PRINTMODE_POSTSCRIPT3_TYPE42             FPDF_PRINTMODE = 7 // to output level 3 PostScript with embedded Type 42 fonts, when applicable, into EMF as a series of GDI comments.
	FPDF_PRINTMODE_POSTSCRIPT3_TYPE42_PASSTHROUGH FPDF_PRINTMODE = 8 // to output level 3 PostScript with embedded Type 42 fonts, when applicable, via ExtEscape() in PASSTHROUGH mode.
)

type FPDF_RENDER_STATUS int

const (
	FPDF_RENDER_STATUS_READY         FPDF_RENDER_STATUS = 0
	FPDF_RENDER_STATUS_TOBECONTINUED FPDF_RENDER_STATUS = 1
	FPDF_RENDER_STATUS_DONE          FPDF_RENDER_STATUS = 2
	FPDF_RENDER_STATUS_FAILED        FPDF_RENDER_STATUS = 3
)

type PDF_FILEAVAIL_LINEARIZATION int

const (
	PDF_FILEAVAIL_LINEARIZATION_UNKNOWN        PDF_FILEAVAIL_LINEARIZATION = -1
	PDF_FILEAVAIL_LINEARIZATION_NOT_LINEARIZED PDF_FILEAVAIL_LINEARIZATION = 0
	PDF_FILEAVAIL_LINEARIZATION_LINEARIZED     PDF_FILEAVAIL_LINEARIZATION = 1
)

type PDF_FILEAVAIL_DATA int

const (
	PDF_FILEAVAIL_DATA_ERROR    PDF_FILEAVAIL_DATA = -1
	PDF_FILEAVAIL_DATA_NOTAVAIL PDF_FILEAVAIL_DATA = 0
	PDF_FILEAVAIL_DATA_AVAIL    PDF_FILEAVAIL_DATA = 1
)

type PDF_FILEAVAIL_FORM int

const (
	PDF_FILEAVAIL_FORM_ERROR    PDF_FILEAVAIL_FORM = -1
	PDF_FILEAVAIL_FORM_NOTAVAIL PDF_FILEAVAIL_FORM = 0
	PDF_FILEAVAIL_FORM_AVAIL    PDF_FILEAVAIL_FORM = 1
	PDF_FILEAVAIL_FORM_NOTEXIST PDF_FILEAVAIL_FORM = 2
)

type PDF_BLEND_MODE string

const (
	PDF_BLEND_MODE_COLOR       PDF_BLEND_MODE = "Color"
	PDF_BLEND_MODE_COLOR_BURN  PDF_BLEND_MODE = "ColorBurn"
	PDF_BLEND_MODE_COLOR_DODGE PDF_BLEND_MODE = "ColorDodge"
	PDF_BLEND_MODE_DARKEN      PDF_BLEND_MODE = "Darken"
	PDF_BLEND_MODE_DIFFERENCE  PDF_BLEND_MODE = "Difference"
	PDF_BLEND_MODE_EXCLUSION   PDF_BLEND_MODE = "Exclusion"
	PDF_BLEND_MODE_HARD_LIGHT  PDF_BLEND_MODE = "HardLight"
	PDF_BLEND_MODE_HUE         PDF_BLEND_MODE = "Hue"
	PDF_BLEND_MODE_LIGHTEN     PDF_BLEND_MODE = "Lighten"
	PDF_BLEND_MODE_LUMINOSITY  PDF_BLEND_MODE = "Luminosity"
	PDF_BLEND_MODE_MULTIPLY    PDF_BLEND_MODE = "Multiply"
	PDF_BLEND_MODE_NORMAL      PDF_BLEND_MODE = "Normal"
	PDF_BLEND_MODE_OVERLAY     PDF_BLEND_MODE = "Overlay"
	PDF_BLEND_MODE_SATURATION  PDF_BLEND_MODE = "Saturation"
	PDF_BLEND_MODE_SCREEN      PDF_BLEND_MODE = "Screen"
	PDF_BLEND_MODE_SOFT_LIGHT  PDF_BLEND_MODE = "SoftLight"
)

type FPDF_FILLMODE int

const (
	FPDF_FILLMODE_NONE      FPDF_FILLMODE = 0
	FPDF_FILLMODE_ALTERNATE FPDF_FILLMODE = 1
	FPDF_FILLMODE_WINDING   FPDF_FILLMODE = 2
)

type FPDF_FONT int

const (
	FPDF_FONT_TYPE1    FPDF_FONT = 1
	FPDF_FONT_TRUETYPE FPDF_FONT = 2
)

type FPDF_PAGEOBJ int

const (
	FPDF_PAGEOBJ_UNKNOWN FPDF_PAGEOBJ = 0
	FPDF_PAGEOBJ_TEXT    FPDF_PAGEOBJ = 1
	FPDF_PAGEOBJ_PATH    FPDF_PAGEOBJ = 2
	FPDF_PAGEOBJ_IMAGE   FPDF_PAGEOBJ = 3
	FPDF_PAGEOBJ_SHADING FPDF_PAGEOBJ = 4
	FPDF_PAGEOBJ_FORM    FPDF_PAGEOBJ = 5
)

type FPDF_LINEJOIN int

const (
	FPDF_LINEJOIN_MITER FPDF_LINEJOIN = 0
	FPDF_LINEJOIN_ROUND FPDF_LINEJOIN = 1
	FPDF_LINEJOIN_BEVEL FPDF_LINEJOIN = 2
)

type FPDF_LINECAP int

const (
	FPDF_LINECAP_BUTT             FPDF_LINECAP = 0
	FPDF_LINECAP_ROUND            FPDF_LINECAP = 1
	FPDF_LINECAP_PROJECTING_SQUAR FPDF_LINECAP = 2
)

type FPDF_SEGMENT int

const (
	FPDF_SEGMENT_UNKNOWN  FPDF_SEGMENT = -1
	FPDF_SEGMENT_LINETO   FPDF_SEGMENT = 0
	FPDF_SEGMENT_BEZIERTO FPDF_SEGMENT = 1
	FPDF_SEGMENT_MOVETO   FPDF_SEGMENT = 2
)

// Refer to PDF Reference version 1.7 table 4.12 for all color space families.
type FPDF_COLORSPACE int

const (
	FPDF_COLORSPACE_UNKNOWN    FPDF_COLORSPACE = 0
	FPDF_COLORSPACE_DEVICEGRAY FPDF_COLORSPACE = 1
	FPDF_COLORSPACE_DEVICERGB  FPDF_COLORSPACE = 2
	FPDF_COLORSPACE_DEVICECMYK FPDF_COLORSPACE = 3
	FPDF_COLORSPACE_CALGRAY    FPDF_COLORSPACE = 4
	FPDF_COLORSPACE_CALRGB     FPDF_COLORSPACE = 5
	FPDF_COLORSPACE_LAB        FPDF_COLORSPACE = 6
	FPDF_COLORSPACE_ICCBASED   FPDF_COLORSPACE = 7
	FPDF_COLORSPACE_SEPARATION FPDF_COLORSPACE = 8
	FPDF_COLORSPACE_DEVICEN    FPDF_COLORSPACE = 9
	FPDF_COLORSPACE_INDEXED    FPDF_COLORSPACE = 10
	FPDF_COLORSPACE_PATTERN    FPDF_COLORSPACE = 11
)

type FPDF_ANNOTATION_SUBTYPE int

const (
	FPDF_ANNOT_SUBTYPE_UNKNOWN        FPDF_ANNOTATION_SUBTYPE = 0
	FPDF_ANNOT_SUBTYPE_TEXT           FPDF_ANNOTATION_SUBTYPE = 1
	FPDF_ANNOT_SUBTYPE_LINK           FPDF_ANNOTATION_SUBTYPE = 2
	FPDF_ANNOT_SUBTYPE_FREETEXT       FPDF_ANNOTATION_SUBTYPE = 3
	FPDF_ANNOT_SUBTYPE_LINE           FPDF_ANNOTATION_SUBTYPE = 4
	FPDF_ANNOT_SUBTYPE_SQUARE         FPDF_ANNOTATION_SUBTYPE = 5
	FPDF_ANNOT_SUBTYPE_CIRCLE         FPDF_ANNOTATION_SUBTYPE = 6
	FPDF_ANNOT_SUBTYPE_POLYGON        FPDF_ANNOTATION_SUBTYPE = 7
	FPDF_ANNOT_SUBTYPE_POLYLINE       FPDF_ANNOTATION_SUBTYPE = 8
	FPDF_ANNOT_SUBTYPE_HIGHLIGHT      FPDF_ANNOTATION_SUBTYPE = 9
	FPDF_ANNOT_SUBTYPE_UNDERLINE      FPDF_ANNOTATION_SUBTYPE = 10
	FPDF_ANNOT_SUBTYPE_SQUIGGLY       FPDF_ANNOTATION_SUBTYPE = 11
	FPDF_ANNOT_SUBTYPE_STRIKEOUT      FPDF_ANNOTATION_SUBTYPE = 12
	FPDF_ANNOT_SUBTYPE_STAMP          FPDF_ANNOTATION_SUBTYPE = 13
	FPDF_ANNOT_SUBTYPE_CARET          FPDF_ANNOTATION_SUBTYPE = 14
	FPDF_ANNOT_SUBTYPE_INK            FPDF_ANNOTATION_SUBTYPE = 15
	FPDF_ANNOT_SUBTYPE_POPUP          FPDF_ANNOTATION_SUBTYPE = 16
	FPDF_ANNOT_SUBTYPE_FILEATTACHMENT FPDF_ANNOTATION_SUBTYPE = 17
	FPDF_ANNOT_SUBTYPE_SOUND          FPDF_ANNOTATION_SUBTYPE = 18
	FPDF_ANNOT_SUBTYPE_MOVIE          FPDF_ANNOTATION_SUBTYPE = 19
	FPDF_ANNOT_SUBTYPE_WIDGET         FPDF_ANNOTATION_SUBTYPE = 20
	FPDF_ANNOT_SUBTYPE_SCREEN         FPDF_ANNOTATION_SUBTYPE = 21
	FPDF_ANNOT_SUBTYPE_PRINTERMARK    FPDF_ANNOTATION_SUBTYPE = 22
	FPDF_ANNOT_SUBTYPE_TRAPNET        FPDF_ANNOTATION_SUBTYPE = 23
	FPDF_ANNOT_SUBTYPE_WATERMARK      FPDF_ANNOTATION_SUBTYPE = 24
	FPDF_ANNOT_SUBTYPE_THREED         FPDF_ANNOTATION_SUBTYPE = 25
	FPDF_ANNOT_SUBTYPE_RICHMEDIA      FPDF_ANNOTATION_SUBTYPE = 26
	FPDF_ANNOT_SUBTYPE_XFAWIDGET      FPDF_ANNOTATION_SUBTYPE = 27
	FPDF_ANNOT_SUBTYPE_REDACT         FPDF_ANNOTATION_SUBTYPE = 28
)

type FPDFANNOT_COLORTYPE int

const (
	FPDFANNOT_COLORTYPE_Color         FPDFANNOT_COLORTYPE = 0
	FPDFANNOT_COLORTYPE_InteriorColor FPDFANNOT_COLORTYPE = 1
)

type FPDF_ANNOT_APPEARANCEMODE int

const (
	FPDF_ANNOT_APPEARANCEMODE_NORMAL   FPDF_ANNOT_APPEARANCEMODE = 0
	FPDF_ANNOT_APPEARANCEMODE_ROLLOVER FPDF_ANNOT_APPEARANCEMODE = 1
	FPDF_ANNOT_APPEARANCEMODE_DOWN     FPDF_ANNOT_APPEARANCEMODE = 2
	FPDF_ANNOT_APPEARANCEMODE_COUNT    FPDF_ANNOT_APPEARANCEMODE = 3
)

type FPDF_FORMFIELD_TYPE int

const (
	FPDF_FORMFIELD_TYPE_UNKNOWN        FPDF_FORMFIELD_TYPE = 0
	FPDF_FORMFIELD_TYPE_PUSHBUTTON     FPDF_FORMFIELD_TYPE = 1
	FPDF_FORMFIELD_TYPE_CHECKBOX       FPDF_FORMFIELD_TYPE = 2
	FPDF_FORMFIELD_TYPE_RADIOBUTTON    FPDF_FORMFIELD_TYPE = 3
	FPDF_FORMFIELD_TYPE_COMBOBOX       FPDF_FORMFIELD_TYPE = 4
	FPDF_FORMFIELD_TYPE_LISTBOX        FPDF_FORMFIELD_TYPE = 5
	FPDF_FORMFIELD_TYPE_TEXTFIELD      FPDF_FORMFIELD_TYPE = 6
	FPDF_FORMFIELD_TYPE_SIGNATURE      FPDF_FORMFIELD_TYPE = 7
	FPDF_FORMFIELD_TYPE_XFA            FPDF_FORMFIELD_TYPE = 8
	FPDF_FORMFIELD_TYPE_XFA_CHECKBOX   FPDF_FORMFIELD_TYPE = 9
	FPDF_FORMFIELD_TYPE_XFA_COMBOBOX   FPDF_FORMFIELD_TYPE = 10
	FPDF_FORMFIELD_TYPE_XFA_IMAGEFIELD FPDF_FORMFIELD_TYPE = 11
	FPDF_FORMFIELD_TYPE_XFA_LISTBOX    FPDF_FORMFIELD_TYPE = 12
	FPDF_FORMFIELD_TYPE_XFA_PUSHBUTTON FPDF_FORMFIELD_TYPE = 13
	FPDF_FORMFIELD_TYPE_XFA_SIGNATURE  FPDF_FORMFIELD_TYPE = 14
	FPDF_FORMFIELD_TYPE_XFA_TEXTFIELD  FPDF_FORMFIELD_TYPE = 15
)

// Refer to PDF Reference (6th edition) table 8.16 for all annotation flags
type FPDF_ANNOT_FLAG int

const (
	FPDF_ANNOT_FLAG_NONE         FPDF_ANNOT_FLAG = 0
	FPDF_ANNOT_FLAG_INVISIBLE    FPDF_ANNOT_FLAG = (1 << 0)
	FPDF_ANNOT_FLAG_HIDDEN       FPDF_ANNOT_FLAG = (1 << 1)
	FPDF_ANNOT_FLAG_PRINT        FPDF_ANNOT_FLAG = (1 << 2)
	FPDF_ANNOT_FLAG_NOZOOM       FPDF_ANNOT_FLAG = (1 << 3)
	FPDF_ANNOT_FLAG_NOROTATE     FPDF_ANNOT_FLAG = (1 << 4)
	FPDF_ANNOT_FLAG_NOVIEW       FPDF_ANNOT_FLAG = (1 << 5)
	FPDF_ANNOT_FLAG_READONLY     FPDF_ANNOT_FLAG = (1 << 6)
	FPDF_ANNOT_FLAG_LOCKED       FPDF_ANNOT_FLAG = (1 << 7)
	FPDF_ANNOT_FLAG_TOGGLENOVIEW FPDF_ANNOT_FLAG = (1 << 8)
)

type FPDF_FORMFLAG int

const (
	FPDF_FORMFLAG_NONE     FPDF_FORMFLAG = 0
	FPDF_FORMFLAG_READONLY FPDF_FORMFLAG = (1 << 0)
	FPDF_FORMFLAG_REQUIRED FPDF_FORMFLAG = (1 << 1)
	FPDF_FORMFLAG_NOEXPORT FPDF_FORMFLAG = (1 << 2)
)

type FXCT int

const (
	FXCT_ARROW FXCT = 0
	FXCT_NESW  FXCT = 1
	FXCT_NWSE  FXCT = 2
	FXCT_VBEAM FXCT = 3
	FXCT_HBEAM FXCT = 4
	FXCT_HAND  FXCT = 5
)

type FPDF_ZOOM_MODE int

const (
	FPDF_ZOOM_MODE_XYZ      FPDF_ZOOM_MODE = 1
	FPDF_ZOOM_MODE_FITPAGE  FPDF_ZOOM_MODE = 2
	FPDF_ZOOM_MODE_FITHORZ  FPDF_ZOOM_MODE = 3
	FPDF_ZOOM_MODE_FITVERT  FPDF_ZOOM_MODE = 4
	FPDF_ZOOM_MODE_FITRECT  FPDF_ZOOM_MODE = 5
	FPDF_ZOOM_MODE_FITBBOX  FPDF_ZOOM_MODE = 6
	FPDF_ZOOM_MODE_FITBHORZ FPDF_ZOOM_MODE = 7
	FPDF_ZOOM_MODE_FITBVERT FPDF_ZOOM_MODE = 8
)

type FPDFDOC_AACTION int

const (
	FPDFDOC_AACTION_WC FPDFDOC_AACTION = 0x10
	FPDFDOC_AACTION_WS FPDFDOC_AACTION = 0x11
	FPDFDOC_AACTION_DS FPDFDOC_AACTION = 0x12
	FPDFDOC_AACTION_WP FPDFDOC_AACTION = 0x13
	FPDFDOC_AACTION_DP FPDFDOC_AACTION = 0x14
)

type FPDFPAGE_AACTION int

const (
	FPDFPAGE_AACTION_OPEN  FPDFPAGE_AACTION = 0 // OPEN (/O) -- An action to be performed when the page is opened
	FPDFPAGE_AACTION_CLOSE FPDFPAGE_AACTION = 1 // CLOSE (/C) -- An action to be performed when the page is closed
)

type FPDF_FORMFIELD int

const (
	FPDF_FORMFIELD_NO_FIELD       FPDF_FORMFIELD = -1
	FPDF_FORMFIELD_UNKNOWN        FPDF_FORMFIELD = 0
	FPDF_FORMFIELD_PUSHBUTTON     FPDF_FORMFIELD = 1
	FPDF_FORMFIELD_CHECKBOX       FPDF_FORMFIELD = 2
	FPDF_FORMFIELD_RADIOBUTTON    FPDF_FORMFIELD = 3
	FPDF_FORMFIELD_COMBOBOX       FPDF_FORMFIELD = 4
	FPDF_FORMFIELD_LISTBOX        FPDF_FORMFIELD = 5
	FPDF_FORMFIELD_TEXTFIELD      FPDF_FORMFIELD = 6
	FPDF_FORMFIELD_SIGNATURE      FPDF_FORMFIELD = 7
	FPDF_FORMFIELD_XFA            FPDF_FORMFIELD = 8
	FPDF_FORMFIELD_XFA_CHECKBOX   FPDF_FORMFIELD = 9
	FPDF_FORMFIELD_XFA_COMBOBOX   FPDF_FORMFIELD = 10
	FPDF_FORMFIELD_XFA_IMAGEFIELD FPDF_FORMFIELD = 11
	FPDF_FORMFIELD_XFA_LISTBOX    FPDF_FORMFIELD = 12
	FPDF_FORMFIELD_XFA_PUSHBUTTON FPDF_FORMFIELD = 13
	FPDF_FORMFIELD_XFA_SIGNATURE  FPDF_FORMFIELD = 14
	FPDF_FORMFIELD_XFA_TEXTFIELD  FPDF_FORMFIELD = 15
)

type FPDF_FORMTYPE int

const (
	FPDF_FORMTYPE_NONE           FPDF_FORMTYPE = 0
	FPDF_FORMTYPE_ACRO_FORM      FPDF_FORMTYPE = 1
	FPDF_FORMTYPE_XFA_FULL       FPDF_FORMTYPE = 2
	FPDF_FORMTYPE_XFA_FOREGROUND FPDF_FORMTYPE = 3
)

type FWL_EVENTFLAG int

const (
	FWL_EVENTFLAG_ShiftKey         FWL_EVENTFLAG = 1 << 0
	FWL_EVENTFLAG_ControlKey       FWL_EVENTFLAG = 1 << 1
	FWL_EVENTFLAG_AltKey           FWL_EVENTFLAG = 1 << 2
	FWL_EVENTFLAG_MetaKey          FWL_EVENTFLAG = 1 << 3
	FWL_EVENTFLAG_KeyPad           FWL_EVENTFLAG = 1 << 4
	FWL_EVENTFLAG_AutoRepeat       FWL_EVENTFLAG = 1 << 5
	FWL_EVENTFLAG_LeftButtonDown   FWL_EVENTFLAG = 1 << 6
	FWL_EVENTFLAG_MiddleButtonDown FWL_EVENTFLAG = 1 << 7
	FWL_EVENTFLAG_RightButtonDown  FWL_EVENTFLAG = 1 << 8
)

type FWL_VKEYCODE int

const (
	FWL_VKEY_Back                      FWL_VKEYCODE = 0x08
	FWL_VKEY_Tab                       FWL_VKEYCODE = 0x09
	FWL_VKEY_NewLine                   FWL_VKEYCODE = 0x0A
	FWL_VKEY_Clear                     FWL_VKEYCODE = 0x0C
	FWL_VKEY_Return                    FWL_VKEYCODE = 0x0D
	FWL_VKEY_Shift                     FWL_VKEYCODE = 0x10
	FWL_VKEY_Control                   FWL_VKEYCODE = 0x11
	FWL_VKEY_Menu                      FWL_VKEYCODE = 0x12
	FWL_VKEY_Pause                     FWL_VKEYCODE = 0x13
	FWL_VKEY_Capital                   FWL_VKEYCODE = 0x14
	FWL_VKEY_Kana                      FWL_VKEYCODE = 0x15
	FWL_VKEY_Hangul                    FWL_VKEYCODE = 0x15
	FWL_VKEY_Junja                     FWL_VKEYCODE = 0x17
	FWL_VKEY_Final                     FWL_VKEYCODE = 0x18
	FWL_VKEY_Hanja                     FWL_VKEYCODE = 0x19
	FWL_VKEY_Kanji                     FWL_VKEYCODE = 0x19
	FWL_VKEY_Escape                    FWL_VKEYCODE = 0x1B
	FWL_VKEY_Convert                   FWL_VKEYCODE = 0x1C
	FWL_VKEY_NonConvert                FWL_VKEYCODE = 0x1D
	FWL_VKEY_Accept                    FWL_VKEYCODE = 0x1E
	FWL_VKEY_ModeChange                FWL_VKEYCODE = 0x1F
	FWL_VKEY_Space                     FWL_VKEYCODE = 0x20
	FWL_VKEY_Prior                     FWL_VKEYCODE = 0x21
	FWL_VKEY_Next                      FWL_VKEYCODE = 0x22
	FWL_VKEY_End                       FWL_VKEYCODE = 0x23
	FWL_VKEY_Home                      FWL_VKEYCODE = 0x24
	FWL_VKEY_Left                      FWL_VKEYCODE = 0x25
	FWL_VKEY_Up                        FWL_VKEYCODE = 0x26
	FWL_VKEY_Right                     FWL_VKEYCODE = 0x27
	FWL_VKEY_Down                      FWL_VKEYCODE = 0x28
	FWL_VKEY_Select                    FWL_VKEYCODE = 0x29
	FWL_VKEY_Print                     FWL_VKEYCODE = 0x2A
	FWL_VKEY_Execute                   FWL_VKEYCODE = 0x2B
	FWL_VKEY_Snapshot                  FWL_VKEYCODE = 0x2C
	FWL_VKEY_Insert                    FWL_VKEYCODE = 0x2D
	FWL_VKEY_Delete                    FWL_VKEYCODE = 0x2E
	FWL_VKEY_Help                      FWL_VKEYCODE = 0x2F
	FWL_VKEY_0                         FWL_VKEYCODE = 0x30
	FWL_VKEY_1                         FWL_VKEYCODE = 0x31
	FWL_VKEY_2                         FWL_VKEYCODE = 0x32
	FWL_VKEY_3                         FWL_VKEYCODE = 0x33
	FWL_VKEY_4                         FWL_VKEYCODE = 0x34
	FWL_VKEY_5                         FWL_VKEYCODE = 0x35
	FWL_VKEY_6                         FWL_VKEYCODE = 0x36
	FWL_VKEY_7                         FWL_VKEYCODE = 0x37
	FWL_VKEY_8                         FWL_VKEYCODE = 0x38
	FWL_VKEY_9                         FWL_VKEYCODE = 0x39
	FWL_VKEY_A                         FWL_VKEYCODE = 0x41
	FWL_VKEY_B                         FWL_VKEYCODE = 0x42
	FWL_VKEY_C                         FWL_VKEYCODE = 0x43
	FWL_VKEY_D                         FWL_VKEYCODE = 0x44
	FWL_VKEY_E                         FWL_VKEYCODE = 0x45
	FWL_VKEY_F                         FWL_VKEYCODE = 0x46
	FWL_VKEY_G                         FWL_VKEYCODE = 0x47
	FWL_VKEY_H                         FWL_VKEYCODE = 0x48
	FWL_VKEY_I                         FWL_VKEYCODE = 0x49
	FWL_VKEY_J                         FWL_VKEYCODE = 0x4A
	FWL_VKEY_K                         FWL_VKEYCODE = 0x4B
	FWL_VKEY_L                         FWL_VKEYCODE = 0x4C
	FWL_VKEY_M                         FWL_VKEYCODE = 0x4D
	FWL_VKEY_N                         FWL_VKEYCODE = 0x4E
	FWL_VKEY_O                         FWL_VKEYCODE = 0x4F
	FWL_VKEY_P                         FWL_VKEYCODE = 0x50
	FWL_VKEY_Q                         FWL_VKEYCODE = 0x51
	FWL_VKEY_R                         FWL_VKEYCODE = 0x52
	FWL_VKEY_S                         FWL_VKEYCODE = 0x53
	FWL_VKEY_T                         FWL_VKEYCODE = 0x54
	FWL_VKEY_U                         FWL_VKEYCODE = 0x55
	FWL_VKEY_V                         FWL_VKEYCODE = 0x56
	FWL_VKEY_W                         FWL_VKEYCODE = 0x57
	FWL_VKEY_X                         FWL_VKEYCODE = 0x58
	FWL_VKEY_Y                         FWL_VKEYCODE = 0x59
	FWL_VKEY_Z                         FWL_VKEYCODE = 0x5A
	FWL_VKEY_LWin                      FWL_VKEYCODE = 0x5B
	FWL_VKEY_Command                   FWL_VKEYCODE = 0x5B
	FWL_VKEY_RWin                      FWL_VKEYCODE = 0x5C
	FWL_VKEY_Apps                      FWL_VKEYCODE = 0x5D
	FWL_VKEY_Sleep                     FWL_VKEYCODE = 0x5F
	FWL_VKEY_NumPad0                   FWL_VKEYCODE = 0x60
	FWL_VKEY_NumPad1                   FWL_VKEYCODE = 0x61
	FWL_VKEY_NumPad2                   FWL_VKEYCODE = 0x62
	FWL_VKEY_NumPad3                   FWL_VKEYCODE = 0x63
	FWL_VKEY_NumPad4                   FWL_VKEYCODE = 0x64
	FWL_VKEY_NumPad5                   FWL_VKEYCODE = 0x65
	FWL_VKEY_NumPad6                   FWL_VKEYCODE = 0x66
	FWL_VKEY_NumPad7                   FWL_VKEYCODE = 0x67
	FWL_VKEY_NumPad8                   FWL_VKEYCODE = 0x68
	FWL_VKEY_NumPad9                   FWL_VKEYCODE = 0x69
	FWL_VKEY_Multiply                  FWL_VKEYCODE = 0x6A
	FWL_VKEY_Add                       FWL_VKEYCODE = 0x6B
	FWL_VKEY_Separator                 FWL_VKEYCODE = 0x6C
	FWL_VKEY_Subtract                  FWL_VKEYCODE = 0x6D
	FWL_VKEY_Decimal                   FWL_VKEYCODE = 0x6E
	FWL_VKEY_Divide                    FWL_VKEYCODE = 0x6F
	FWL_VKEY_F1                        FWL_VKEYCODE = 0x70
	FWL_VKEY_F2                        FWL_VKEYCODE = 0x71
	FWL_VKEY_F3                        FWL_VKEYCODE = 0x72
	FWL_VKEY_F4                        FWL_VKEYCODE = 0x73
	FWL_VKEY_F5                        FWL_VKEYCODE = 0x74
	FWL_VKEY_F6                        FWL_VKEYCODE = 0x75
	FWL_VKEY_F7                        FWL_VKEYCODE = 0x76
	FWL_VKEY_F8                        FWL_VKEYCODE = 0x77
	FWL_VKEY_F9                        FWL_VKEYCODE = 0x78
	FWL_VKEY_F10                       FWL_VKEYCODE = 0x79
	FWL_VKEY_F11                       FWL_VKEYCODE = 0x7A
	FWL_VKEY_F12                       FWL_VKEYCODE = 0x7B
	FWL_VKEY_F13                       FWL_VKEYCODE = 0x7C
	FWL_VKEY_F14                       FWL_VKEYCODE = 0x7D
	FWL_VKEY_F15                       FWL_VKEYCODE = 0x7E
	FWL_VKEY_F16                       FWL_VKEYCODE = 0x7F
	FWL_VKEY_F17                       FWL_VKEYCODE = 0x80
	FWL_VKEY_F18                       FWL_VKEYCODE = 0x81
	FWL_VKEY_F19                       FWL_VKEYCODE = 0x82
	FWL_VKEY_F20                       FWL_VKEYCODE = 0x83
	FWL_VKEY_F21                       FWL_VKEYCODE = 0x84
	FWL_VKEY_F22                       FWL_VKEYCODE = 0x85
	FWL_VKEY_F23                       FWL_VKEYCODE = 0x86
	FWL_VKEY_F24                       FWL_VKEYCODE = 0x87
	FWL_VKEY_NunLock                   FWL_VKEYCODE = 0x90
	FWL_VKEY_Scroll                    FWL_VKEYCODE = 0x91
	FWL_VKEY_LShift                    FWL_VKEYCODE = 0xA0
	FWL_VKEY_RShift                    FWL_VKEYCODE = 0xA1
	FWL_VKEY_LControl                  FWL_VKEYCODE = 0xA2
	FWL_VKEY_RControl                  FWL_VKEYCODE = 0xA3
	FWL_VKEY_LMenu                     FWL_VKEYCODE = 0xA4
	FWL_VKEY_RMenu                     FWL_VKEYCODE = 0xA5
	FWL_VKEY_BROWSER_Back              FWL_VKEYCODE = 0xA6
	FWL_VKEY_BROWSER_Forward           FWL_VKEYCODE = 0xA7
	FWL_VKEY_BROWSER_Refresh           FWL_VKEYCODE = 0xA8
	FWL_VKEY_BROWSER_Stop              FWL_VKEYCODE = 0xA9
	FWL_VKEY_BROWSER_Search            FWL_VKEYCODE = 0xAA
	FWL_VKEY_BROWSER_Favorites         FWL_VKEYCODE = 0xAB
	FWL_VKEY_BROWSER_Home              FWL_VKEYCODE = 0xAC
	FWL_VKEY_VOLUME_Mute               FWL_VKEYCODE = 0xAD
	FWL_VKEY_VOLUME_Down               FWL_VKEYCODE = 0xAE
	FWL_VKEY_VOLUME_Up                 FWL_VKEYCODE = 0xAF
	FWL_VKEY_MEDIA_NEXT_Track          FWL_VKEYCODE = 0xB0
	FWL_VKEY_MEDIA_PREV_Track          FWL_VKEYCODE = 0xB1
	FWL_VKEY_MEDIA_Stop                FWL_VKEYCODE = 0xB2
	FWL_VKEY_MEDIA_PLAY_Pause          FWL_VKEYCODE = 0xB3
	FWL_VKEY_MEDIA_LAUNCH_Mail         FWL_VKEYCODE = 0xB4
	FWL_VKEY_MEDIA_LAUNCH_MEDIA_Select FWL_VKEYCODE = 0xB5
	FWL_VKEY_MEDIA_LAUNCH_APP1         FWL_VKEYCODE = 0xB6
	FWL_VKEY_MEDIA_LAUNCH_APP2         FWL_VKEYCODE = 0xB7
	FWL_VKEY_OEM_1                     FWL_VKEYCODE = 0xBA
	FWL_VKEY_OEM_Plus                  FWL_VKEYCODE = 0xBB
	FWL_VKEY_OEM_Comma                 FWL_VKEYCODE = 0xBC
	FWL_VKEY_OEM_Minus                 FWL_VKEYCODE = 0xBD
	FWL_VKEY_OEM_Period                FWL_VKEYCODE = 0xBE
	FWL_VKEY_OEM_2                     FWL_VKEYCODE = 0xBF
	FWL_VKEY_OEM_3                     FWL_VKEYCODE = 0xC0
	FWL_VKEY_OEM_4                     FWL_VKEYCODE = 0xDB
	FWL_VKEY_OEM_5                     FWL_VKEYCODE = 0xDC
	FWL_VKEY_OEM_6                     FWL_VKEYCODE = 0xDD
	FWL_VKEY_OEM_7                     FWL_VKEYCODE = 0xDE
	FWL_VKEY_OEM_8                     FWL_VKEYCODE = 0xDF
	FWL_VKEY_OEM_102                   FWL_VKEYCODE = 0xE2
	FWL_VKEY_ProcessKey                FWL_VKEYCODE = 0xE5
	FWL_VKEY_Packet                    FWL_VKEYCODE = 0xE7
	FWL_VKEY_Attn                      FWL_VKEYCODE = 0xF6
	FWL_VKEY_Crsel                     FWL_VKEYCODE = 0xF7
	FWL_VKEY_Exsel                     FWL_VKEYCODE = 0xF8
	FWL_VKEY_Ereof                     FWL_VKEYCODE = 0xF9
	FWL_VKEY_Play                      FWL_VKEYCODE = 0xFA
	FWL_VKEY_Zoom                      FWL_VKEYCODE = 0xFB
	FWL_VKEY_NoName                    FWL_VKEYCODE = 0xFC
	FWL_VKEY_PA1                       FWL_VKEYCODE = 0xFD
	FWL_VKEY_OEM_Clear                 FWL_VKEYCODE = 0xFE
	FWL_VKEY_Unknown                   FWL_VKEYCODE = 0
)

type FPDF_ANNOT_AACTION int

const (
	FPDF_ANNOT_AACTION_KEY_STROKE FPDF_ANNOT_AACTION = 12
	FPDF_ANNOT_AACTION_FORMAT     FPDF_ANNOT_AACTION = 13
	FPDF_ANNOT_AACTION_VALIDATE   FPDF_ANNOT_AACTION = 14
	FPDF_ANNOT_AACTION_CALCULATE  FPDF_ANNOT_AACTION = 15
)
