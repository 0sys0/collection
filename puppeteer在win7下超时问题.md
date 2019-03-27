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
```
const puppeteer = require('puppeteer');

async function getPic() {
  const browser = await puppeteer.launch({args: [
    '--proxy-server="direct://"',
    '--proxy-bypass-list=*'
  ]});
  const page = await browser.newPage();
  await page.goto('https://www.google.com');
  await page.screenshot({path: 'google.png'});

  await browser.close();
}

getPic();
```
因此需要在使用reuqests_html或pyppeteer时要这样修改：
```
#pyppeteer/laucher.py, line 102
        if 'headless' not in self.options or self.options.get('headless'):
            self.chrome_args.extend([
                '--headless',
                '--disable-gpu',
                '--hide-scrollbars',
                '--mute-audio',
                #windows系统专用
                '--proxy-server="direct://"',
                '--proxy-bypass-list=*'
            ])
```
