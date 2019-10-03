function gorender(input,sel){
        input="var el = "+input;
        eval(Babel.transform(input, { presets: ['es2015','react'], "sourceType": "script" }).code);
        ReactDOM.render(el,document.querySelector(sel));
}
