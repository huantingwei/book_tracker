module.exports = {
    root: true,
    env: {
        browser: true,
        node: true
    },
    parserOptions: {
        parser: 'babel-eslint',
        ecmaVersion: 2018,
        sourceType: 'module'
    },
    extends: [
        "eslint:recommended",
        "plugin:react/recommended",
        "plugin:react-app/recommended",
        'plugin:prettier/recommended'
    ],
    plugins: ['react', 'react-hooks'],
    // add your custom rules here
    rules: {
        "react-app/react/react-in-jsx-scope": ["warn"],
        'react/prop-types': 1,
        "react/jsx-uses-vars": "error",
        'react-hooks/rules-of-hooks': 'error', // 檢查 Hook 的規則
        'react-hooks/exhaustive-deps': 'warn', // 檢查 effect 的相依性
    },
    
}
