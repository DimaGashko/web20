module.exports = {
    env: {
        'browser': true,
        'node': true,
        'es6': true,
    },
    extends: [
        'eslint:recommended',
    ],
    rules: {
        'semi': ['warn', 'always'],
        'comma-dangle': ['warn', 'always-multiline'],
        'quotes': ['warn', 'single', {
            'allowTemplateLiterals': true,
            'avoidEscape': true,
        }],

        'no-unused-vars': ['off'],
        'no-unreachable': ['off'],
        'no-constant-condition': ['off'],
        'no-useless-escape': ['off'],
    },
    parser: 'babel-eslint',
    parserOptions: {
        ecmaVersion: 2020,
        sourceType: 'module',
    },
};