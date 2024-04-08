# PDS - Parser and Builder

The PDS (Parser and Builder) system is designed for encoding and decoding structured data within Field 48 of Clearing-Acquirer financial transactions. This field often contains complex and variable-length data critical for transaction processing. The PDS system employs a Tag-Length-Value (TLV) encoding scheme to efficiently manage this data.

## PDS Encoding Format

PDS data elements are encoded in a Tag-Length-Data (TLD) format, facilitating flexible and dynamic parsing. Each data element within a PDS is structured as follows:

- **Tag**: A unique identifier for the data element (ID).
- **Length**: The length of the data element, allowing the parser to determine how many characters to read for the value.
- **Data**: The actual data of the element, which can vary in type and format.

### Structure Overview

| Value Number | Positions | Description                |
|--------------|-----------|----------------------------|
| 1            | 1-4       | First PDS tag (ID)         |
| 2            | 5-7       | First PDS data length      |
| 3            | 8-999     | First PDS data             |

Subfields 1-3 must be repeated for each PDS element until all data within Field 48 is fully represented. The parser must be capable of dynamically adjusting to the length and number of PDS elements within a given transaction message.

## Parsing Field 48

Field 48 in Clearing-Acquirer transactions is a critical component containing detailed transaction information. The PDS system allows for the structured and efficient parsing of this data, ensuring that all necessary information is accurately extracted and processed.

The parser must iterate through the Field 48 data, identifying each PDS based on its tag, determining the data length, and then extracting the corresponding data value. This process is repeated until all PDS elements within Field 48 are parsed.

### Example

Consider a Field 48 value structured as follows:

`00011002345678901230005ABCDE`

- **Tag**: `0001` - Indicates a specific data element type.
- **Length**: `002` - The length of the data, indicating that 2 characters should be read.
- **Data**: `34` - The data value for this element.
- The process repeats for the next element, with a new tag, length, and data value.

## Conclusion

The PDS system provides a robust framework for encoding and decoding complex data structures within financial transactions, specifically within the context of Clearing-Acquirer transactions' Field 48. By employing a TLV encoding scheme, the PDS ensures data integrity and flexibility in transaction processing.
