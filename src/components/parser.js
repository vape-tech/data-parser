const fs = require('fs');
const path = require('path');
const { parse } = require('csv-parse/sync');

class DataParser {
    constructor(filePath) {
        this.filePath = path.resolve(__dirname, filePath);
    }

    parseCSV() {
        try {
            const fileContent = fs.readFileSync(this.filePath, 'utf8');
            const records = parse(fileContent, {
                columns: true,
                skip_empty_lines: true,
            });
            return records;
        } catch (error) {
            console.error(`Error parsing CSV file: ${error.message}`);
            throw error;
        }
    }

    saveAsJSON(outputPath, data) {
        try {
            const jsonContent = JSON.stringify(data, null, 2);
            fs.writeFileSync(outputPath, jsonContent, 'utf8');
            console.log(`File saved successfully at ${outputPath}`);
        } catch (error) {
            console.error(`Error saving JSON file: ${error.message}`);
            throw error;
        }
    }
}

module.exports = DataParser;