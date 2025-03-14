package filesystem

import (
	"os"
	"server/api/dj/log"
	"strings"

	"github.com/dhowden/tag"
)

// cspell:disable
// AENC - Audio encryption
// APIC - Attached picture
// ASPI - Audio seek point index
// COMM - Comments
// COMR - Commercial frame
// ENCR - Encryption method registration
// EQU2 - Equalisation (2)
// ETCO - Event timing codes
// GEOB - General encapsulated object
// GRID - Group identification registration
// LINK - Linked information
// MCDI - Music CD identifier
// MLLT - MPEG location lookup table
// OWNE - Ownership frame
// PRIV - Private frame
// PCNT - Play counter
// POPM - Popularimeter
// POSS - Position synchronisation frame
// RBUF - Recommended buffer size
// RVA2 - Relative volume adjustment (2)
// RVRB - Reverb
// SEEK - Seek frame
// SIGN - Signature frame
// SYLT - Synchronised lyric/text
// SYTC - Synchronised tempo codes
// TALB - Album/Movie/Show title
// TBPM - BPM (beats per minute)
// TCOM - Composer
// TCON - Content type
// TCOP - Copyright message
// TDEN - Encoding time
// TDLY - Playlist delay
// TDOR - Original release time
// TDRC - Recording time
// TDRL - Release time
// TDTG - Tagging time
// TENC - Encoded by
// TEXT - Lyricist/Text writer
// TFLT - File type
// TIPL - Involved people list
// TIT1 - Content group description
// TIT2 - Title/songname/content description
// TIT3 - Subtitle/Description refinement
// TKEY - Initial key
// TLAN - Language(s)
// TLEN - Length
// TMCL - Musician credits list
// TMED - Media type
// TMOO - Mood
// TOAL - Original album/movie/show title
// TOFN - Original filename
// TOLY - Original lyricist(s)/text writer(s)
// TOPE - Original artist(s)/performer(s)
// TOWN - File owner/licensee
// TPE1 - Lead performer(s)/Soloist(s)
// TPE2 - Band/orchestra/accompaniment
// TPE3 - Conductor/performer refinement
// TPE4 - Interpreted, remixed, or otherwise modified by
// TPOS - Part of a set
// TPRO - Produced notice
// TPUB - Publisher
// TRCK - Track number/Position in set
// TRSN - Internet radio station name
// TRSO - Internet radio station owner
// TSOA - Album sort order
// TSOP - Performer sort order
// TSOT - Title sort order
// TSRC - ISRC (international standard recording code)
// TSSE - Software/Hardware and settings used for encoding
// TSST - Set subtitle
// TXXX - User defined text information frame
// UFID - Unique file identifier
// USER - Terms of use
// USLT - Unsynchronised lyric/text transcription
// WCOM - Commercial information
// WCOP - Copyright/Legal information
// WOAF - Official audio file webpage
// WOAR - Official artist/performer webpage
// WOAS - Official audio source webpage
// WORS - Official Internet radio station homepage
// WPAY - Payment
// WPUB - Publishers official webpage
// WXXX - User defined URL link frame
// cspell:enable

func ReadMeta(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Log(err)
		return nil, err
	}
	defer file.Close()

	tags, err := tag.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	meta := tags.Raw()
	for k, v := range meta {
		// trim images
		if strings.HasPrefix(k, "APIC") {
			v.(*tag.Picture).Data = []byte{0, 1, 2, 3}
		}
	}

	return meta, nil
}
