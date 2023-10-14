type Var = {
    name: string,
    type: string 
};

export type Fn = {
    args: Var[],
    return: string
};

type Category = 
    "Arrays & Hashing" | 
    "Two Pointers" | 
    "Stack" | 
    "Binar Search" | 
    "Sliding Window" | 
    "Linked List" | 
    "Trees" | 
    "Tries" | 
    "Backtracking" |
    "Heap / Priority Queue" | 
    "Graph" | 
    "1-D DP" | 
    "Intervals" | 
    "Greedy" | 
    "Advanced Graphs" | 
    "2-D DP" | 
    "Bit Manipulation" | 
    "Math & Geometry";

export const categories = [
    "Arrays & Hashing", 
    "Two Pointers", 
    "Stack", 
    "Binar Search", 
    "Sliding Window", 
    "Linked List", 
    "Trees", 
    "Tries", 
    "Backtracking",
    "Heap / Priority Queue", 
    "Graph", 
    "1-D DP", 
    "Intervals", 
    "Greedy", 
    "Advanced Graphs", 
    "2-D DP", 
    "Bit Manipulation", 
    "Math & Geometry"
];

export type Class = {
    member: Var[],
    methods: Fn[]
};

type Difficulty = "easy" | "medium" | "hard";

export const difficulties = [
    "easy" , "medium" , "hard"
];

export type Question = {
    name: string
    category: Category,
    difficulty: Difficulty,
    type: Fn | Class 
}

export const questions: Question[] = [
    {
       name: "containsDuplicate",
       category: "Arrays & Hashing",
       difficulty: "easy",
       type: {
           args: [
               { name: "nums", type: "int[]" }
           ],
           return: "bool"
       }
    },
    {
       name: "isAnagram",
       category: "Arrays & Hashing",
       difficulty: "easy",
       type: {
           args: [
               { name: "s", type: "str" },
               { name: "t", type: "str" }
           ],
           return: "bool"
       }
    },
    {
        name: "twoSum",
        category: "Arrays & Hashing",
        difficulty: "easy",
        type: {
            args: [
                { name: "nums", type: "int[]" },
                { name: "target", type: "int" }
            ],
            return: "int[]"
        }
    },
    {
        name: "groupAnagrmas",
        category: "Arrays & Hashing",
        difficulty: "medium",
        type: {
            args: [
                { name: "strs", type: "str[]" },
            ],
            return: "str[][]"
        }
    }
];
