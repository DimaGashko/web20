import showdown from 'showdown';

const mdConverter = new showdown.Converter({
   openLinksInNewWindow: false,
});

const html = mdConverter.makeHtml(`

# Web 2.0

Lorem ipsum dolor sit amet consectetur adipisicing elit. Molestiae amet iure deleniti tempore incidunt iusto maxime, 
quaerat laboriosam tenetur similique pariatur adipisci ratione. Assumenda, quo dicta laborum beatae error quasi!

## Sub Title Here
<script>alert('This is one');</script>
[some text](javascript:alert('xss'))
Lorem ipsum dolor sit amet consectetur adipisicing elit. Molestiae amet iure deleniti tempore incidunt iusto maxime, 
quaerat laboriosam tenetur similique pariatur adipisci ratione. Assumenda, quo dicta laborum beatae error quasi!Lorem
ipsum dolor sit amet consectetur adipisicing elit. Molestiae amet iure deleniti tempore incidunt iusto maxime
laboriosam tenetur similique pariatur adipisci ratione. Assumenda, quo dicta laborum beatae error quasi!

- Item 1
- Item 2

1. Item 1
1. Item 2

`);

//document.body.innerHTML = html;
//console.log(mdConverter.getMetadata());