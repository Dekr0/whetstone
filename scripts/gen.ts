import { mkdir } from "fs/promises";
import { Fn, Class, questions, Question, categories, difficulties } from "./questions";
import { join } from "path";

function isFn(t: Fn | Class): t is Fn {
    return (t as Fn).return !== undefined;
}

const tsTypeTranslate = function(t: string) {
    const b = t.indexOf('[]');
    let before = b !== -1 ? t.slice(0, b) : t;
    let after: string;
    switch (before) {
        case "int": {
            after = "number";
            break;
        }
        case "str": {
            after = "string";
            break;
        }
        case "bool": {
            after = "boolean";
            break;
        }
        default: {
            after = before;
            console.log(`potential user-defined type / class / interface ${after}`);
            break;
        }
    }

    return t.replace(before, after);
};

function genFn(question: Question) {
    let code = "";
    code += `function ${question.name}(`;
    const fn = question.type as Fn;
    for (let i = 0; i < fn.args.length; i++) {
        code += `${fn.args[i].name}:${tsTypeTranslate(fn.args[i].type)}`;
        if (i < fn.args.length - 1)
            code += ",";
        else
            code += ")";
    }
    code += `:${tsTypeTranslate(fn.return)}{}`;
}

function genClass(question: Question) {
    let code = "";
    
    return code;
}

function ts() {
    for (const question of questions) {
        let code = isFn(question.type) ? genFn(question) : genClass(question);
    }
}

function gen() {
    mkdir("day", { recursive: false }).catch(_ => { 
            throw new Error(`Failed to create directory "day`);
    });

    for (const category of categories) {
        const c = join("day", category);
        mkdir(c, { recursive: false }).catch(_ => {
                throw new Error(`Failed to create directory ${c}`);
        });
        for (const difficulty of difficulties) {
            const d = join(c, difficulty);
            mkdir(d, { recursive: false }).catch(_ => {
                throw new Error(`Failed to create directory ${d}`);
            });
        }
    }
}

gen();
