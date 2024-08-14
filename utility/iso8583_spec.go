package utility

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
)

var Spec87RK *iso8583.MessageSpec = &iso8583.MessageSpec{
	Name: "ISO 8583 v1987 ASCII",
	Fields: map[int]field.Field{
		0: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message Type Indicator",
			Enc:         encoding.BCD,
			Pref:        prefix.BCD.Fixed,
		}),
		1: field.NewBitmap(&field.Spec{
			Length:      8,
			Description: "Bitmap",
			Enc:         encoding.Binary,
			Pref:        prefix.Binary.Fixed,
		}),
		2: field.NewString(&field.Spec{
			Length:      19,
			Description: "Primary Account Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		3: field.NewString(&field.Spec{
			Length:      6,
			Description: "Processing Code",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		4: field.NewString(&field.Spec{
			Length:      12,
			Description: "Transaction Amount",
			Enc:         encoding.Binary,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		5: field.NewString(&field.Spec{
			Length:      12,
			Description: "Settlement Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		6: field.NewString(&field.Spec{
			Length:      12,
			Description: "Billing Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		7: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transmission Date & Time",
			Enc:         encoding.Binary,
			Pref:        prefix.ASCII.Fixed,
		}),
		8: field.NewString(&field.Spec{
			Length:      8,
			Description: "Billing Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		9: field.NewString(&field.Spec{
			Length:      8,
			Description: "Settlement Conversion Rate",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		10: field.NewString(&field.Spec{
			Length:      8,
			Description: "Cardholder Billing Conversion Rate",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		11: field.NewString(&field.Spec{
			Length:      6,
			Description: "Systems Trace Audit Number (STAN)",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
		}),
		12: field.NewString(&field.Spec{
			Length:      6,
			Description: "Local Transaction Time",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
		}),
		13: field.NewString(&field.Spec{
			Length:      4,
			Description: "Local Transaction Date",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
		}),
		14: field.NewString(&field.Spec{
			Length:      4,
			Description: "Expiration Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		15: field.NewString(&field.Spec{
			Length:      4,
			Description: "Settlement Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		16: field.NewString(&field.Spec{
			Length:      4,
			Description: "Currency Conversion Date",
			Enc:         encoding.ASCIIHexToBytes,
			Pref:        prefix.ASCII.Fixed,
		}),
		17: field.NewString(&field.Spec{
			Length:      4,
			Description: "Capture Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		18: field.NewString(&field.Spec{
			Length:      4,
			Description: "Merchant Type",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		19: field.NewString(&field.Spec{
			Length:      3,
			Description: "Acquiring Institution Country Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		20: field.NewString(&field.Spec{
			Length:      3,
			Description: "PAN Extended Country Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		21: field.NewString(&field.Spec{
			Length:      3,
			Description: "Forwarding Institution Country Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		22: field.NewString(&field.Spec{
			Length:      3,
			Description: "Point of Sale (POS) Entry Mode",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
		}),
		23: field.NewString(&field.Spec{
			Length:      3,
			Description: "Card Sequence Number (CSN)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		24: field.NewString(&field.Spec{
			Length:      3,
			Description: "Function Code",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		25: field.NewString(&field.Spec{
			Length:      2,
			Description: "Point of Service Condition Code",
			Enc:         encoding.BCD,
			Pref:        prefix.ASCII.Fixed,
		}),
		26: field.NewString(&field.Spec{
			Length:      2,
			Description: "Point of Service PIN Capture Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		27: field.NewString(&field.Spec{
			Length:      1,
			Description: "Authorizing Identification Response Length",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		28: field.NewString(&field.Spec{
			Length:      9,
			Description: "Transaction Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		29: field.NewString(&field.Spec{
			Length:      9,
			Description: "Settlement Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		30: field.NewString(&field.Spec{
			Length:      9,
			Description: "Transaction Processing Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		31: field.NewString(&field.Spec{
			Length:      9,
			Description: "Settlement Processing Fee Amount",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		32: field.NewString(&field.Spec{
			Length:      11,
			Description: "Acquiring Institution Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		33: field.NewString(&field.Spec{
			Length:      11,
			Description: "Forwarding Institution Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.Binary.LL,
		}),
		34: field.NewString(&field.Spec{
			Length:      28,
			Description: "Extended Primary Account Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		35: field.NewString(&field.Spec{
			Length:      37,
			Description: "Track 2 Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		36: field.NewString(&field.Spec{
			Length:      104,
			Description: "Track 3 Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		37: field.NewString(&field.Spec{
			Length:      12,
			Description: "Retrieval Reference Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		38: field.NewString(&field.Spec{
			Length:      6,
			Description: "Authorization Identification Response",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		39: field.NewString(&field.Spec{
			Length:      2,
			Description: "Response Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		40: field.NewString(&field.Spec{
			Length:      3,
			Description: "Service Restriction Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		41: field.NewString(&field.Spec{
			Length:      8,
			Description: "Card Acceptor Terminal Identification",
			Enc:         encoding.Binary,
			Pref:        prefix.ASCII.Fixed,
		}),
		42: field.NewString(&field.Spec{
			Length:      15,
			Description: "Card Acceptor Identification Code",
			Enc:         encoding.Binary,
			Pref:        prefix.ASCII.Fixed,
		}),
		43: field.NewString(&field.Spec{
			Length:      40,
			Description: "Card Acceptor Name/Location",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		44: field.NewString(&field.Spec{
			Length:      99,
			Description: "Additional Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		45: field.NewString(&field.Spec{
			Length:      76,
			Description: "Track 1 Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		46: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data (ISO)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		47: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		48: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional data (Private)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		49: field.NewString(&field.Spec{
			Length:      3,
			Description: "Transaction Currency Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		50: field.NewString(&field.Spec{
			Length:      3,
			Description: "Settlement Currency Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		51: field.NewString(&field.Spec{
			Length:      3,
			Description: "Cardholder Billing Currency Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		52: field.NewString(&field.Spec{
			Length:      8,
			Description: "PIN Data",
			Enc:         encoding.ASCIIHexToBytes,
			Pref:        prefix.ASCII.Fixed,
		}),
		53: field.NewString(&field.Spec{
			Length:      16,
			Description: "Security Related Control Information",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		54: field.NewString(&field.Spec{
			Length:      120,
			Description: "Additional Amounts",
			Enc:         encoding.ASCII,
			Pref:        prefix.BCD.LLL,
		}),
		55: field.NewString(&field.Spec{
			Length:      999,
			Description: "ICC Data – EMV Having Multiple Tags",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		56: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (ISO)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		57: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.Binary,
			Pref:        prefix.BCD.LLL,
		}),
		58: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.Binary,
			Pref:        prefix.BCD.LLL,
		}),
		59: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		60: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (National)",
			Enc:         encoding.Binary,
			Pref:        prefix.BCD.LLL,
		}),
		61: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (Private)",
			Enc:         encoding.Binary,
			Pref:        prefix.BCD.LLL,
		}),
		62: field.NewString(&field.Spec{
			Length:      999,
			Description: "Reserved (Private)",
			Enc:         encoding.Binary,
			Pref:        prefix.BCD.LLL,
		}),
		63: field.NewString(&field.Spec{
			Length:      99,
			Description: "Reserved (Private)",
			Enc:         encoding.Binary,
			Pref:        prefix.BCD.LLL,
		}),
		64: field.NewString(&field.Spec{
			Length:      8,
			Description: "Message Authentication Code (MAC)",
			Enc:         encoding.ASCIIHexToBytes,
			Pref:        prefix.ASCII.Fixed,
		}),
		90: field.NewString(&field.Spec{
			Length:      42,
			Description: "Original Data Elements",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
	},
}
