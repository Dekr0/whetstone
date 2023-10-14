import { outro, select } from '@clack/prompts';
import { log } from 'console';
import { existsSync, readdirSync, writeFile } from 'fs';
import { copyFile, mkdir, readFile } from 'fs/promises';
import { dirname, join } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const src = join(__dirname, '..', 'src');
const practice = join(src, 'practice');
const template = join(src, 'template');
const status = join(__dirname, 'status.json');
const langs = readdirSync(template);

const labels = {
    'ts': 'TypeScript',
    'js': 'JavaScript',
    'rust': 'Rust',
    'go': 'Go'
}

async function copyDir(src, dest) {
    await mkdir(dest, { recursive: true });

    let entries = readdirSync(src, { withFileTypes : true });
    for (let entry of entries) {
        let s = join(src, entry.name);
        let d = join(dest, entry.name);

        entry.isDirectory() ? await copyDir(s, d) : await copyFile(s, d);
    }
}

async function gen() {
    let history = {};
    if (!existsSync(status)) {
        for (const lang of langs) {
            history[lang] = 1;
        }

        writeFile(status, JSON.stringify(history), 'utf8', err => {
            if (err) throw err;
        });
    } else {
        history = JSON.parse(await readFile(status));
    }

    const value = await select({
        message: 'Pick a language for template generation',
        options: [
            { value: 'ts', label: 'TypeScript' },
            { value: 'js', label: 'JavaScript' },
            { value: 'rust', label: 'Rust' },
            { value: 'go', label: 'Go' }
        ],
    });

    if (langs.includes(value)) {
        const dest = join(practice, `try${history[value]++}`)
        try {
            await copyDir(join(template, value), dest);

            writeFile(status, JSON.stringify(history), 'utf8', err => {
                if (err) throw err;
            });
        } catch (err) {
            log(err.message);
        }
    } else {
        outro(`${labels[value]} is under construction`);
    }
}

gen();
