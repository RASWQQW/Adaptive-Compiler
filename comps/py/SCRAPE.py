import mechanize as mec
import bs4
url = 'https://thepythoncode.com/assistant/code-converter/c++/python/'
url2 = 'https://rextester.com/'
url3 = 'https://rextester.com/l/python3_online_compiler'
br = mec.Browser()
br.open(url3)

print(br.geturl())


for ff in br.forms():
    print(ff)

formIs = br.select_form(nr=0)
print(br.form)
toclick = br.form.find_control(nr=5)
print('get control', toclick.value)


chance = br._forms[0].controls[5]
print('5th cont ', chance.attrs['value'])#(nr=0).controls(nr=5).value)
#res = br.click(nr=5)
br.click(nr=6)

# for vv in br.form.attrs:
#     print('Attr', vv)
# print(br.form._labels)

br.form['Program'] = 'print("Apple")'
print(br.form['Program'])
print(br.form.attrs)
res1 = br.form.find_control(id='Run')
print('ResVal:', res1)

#print(mec.urlopen(br).read())

# bsv = bs4.BeautifulSoup(res1, 'html.parser')
# print(bsv.find('pre', id='Result').text)
#dd = br.form.find_control(id='Program-code') # type='textarea',
#dd.value = 'Apple'
#print(dd.value)

# for dd in br.form._forms:
#     print(dd)
# for vv in  br.forms():
#     print(vv)

#br['TextControl'] = 'std:cout<<"Apple";'
