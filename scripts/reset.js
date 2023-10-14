import { confirm } from "@clack/prompts";
import { existsSync, writeFile } from "fs";
import { readFile, rm } from "fs/promises";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const src = join(__dirname, '..', 'src');
const practice = join(src, 'practice');
const status = join(__dirname, 'status.json');

async function reset() {

    const c = await confirm({
        message: 'Reset all progress?',
    });

    if (!c) {
        return;
    }

    const guard = await confirm({
        message: 'Warning: all progress will reset',
    });

    if (!guard) {
        return;
    }

    await rm(practice, { recursive: true });

    let history = {};
    if (existsSync(status)) {
        history = JSON.parse(await readFile(status));
    }

    for (const prop in history) {
        history[prop] = 1;
    }

    writeFile(status, JSON.stringify(history), 'utf8', err => {
        if (err) throw err;
    });
}

reset();
