puppeteer在win7下超时问题
> https://github.com/GoogleChrome/puppeteer/issues/2391
```
const browser = await puppeteer.launch({
  headless: true,
  args: [
    '--proxy-server="direct://"',
    '--proxy-bypass-list=*'
  ]
});
```
