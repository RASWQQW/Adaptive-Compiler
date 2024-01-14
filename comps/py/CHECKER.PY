# import os
#
#
# text = '''def apple(ss):\n\tprint(ss)\n\tapple(exer)
# '''
# print(text)
# dd = 12
# exer = 'HELLO DANG'
# print(exec(compile(text, 'test', 'exec')))
# #apple(dd)


import telethon as th
import asyncio



api_id = 16147296
api_hash = "b37c46483b8f55c3a0221e9c3b859825"
botApi = "6059490532:AAHk7sYVhcRhBg7Opfn3ieae6_wo7hfWtNU"

import requests

# response = requests.post(
#     url='https://api.telegram.org/bot{0}/{1}'.format(botApi, 'sendMessage'),
#     data={'chat_id': 5720028884, 'text': 'hello friend'}
# ).json()
# print(response)

import requests
url = "https://api.jdoodle.com/v1/execute"
api_key = "b86eafbab040cfd06f729b4f7d233f2d"
code = "print('Hello, JDoodle!')"
language = "python3"
payload = {
    "clientId": api_key,
    "clientSecret": '85abad0125e7efdfeb9db2deb1b4155b919c65963deb43a0391ee6209cc71ba9',
    "script": code,
    "language": language,
    }
response = requests.post(url, json=payload)
print(response.json())

async def runner(lang: str, code: str):
    cl = th.TelegramClient('Checker', api_id, api_hash)
    cl.connect()

    mainEnt = 5720028884
    chat = -1001862349713
    # 2872972703456

    async with cl as con:
        #async for val in con.iter_dialogs():
            # if val.name.__contains__('ChatGPT'):
            #     print(val)

        pureCode = 'without commentary'
        text = 'Convert from ' + lang + ' to Python ' + pureCode + '\n\n' + code
        vals = await con.get_entity(mainEnt)
        print(vals)
        ms = await con.send_message(vals, text)

        print('end')
        while True:
            await asyncio.sleep(0.140)
            valsn = await con.get_messages(vals, limit=1)
            if valsn[0].reply_to.reply_to_msg_id == ms.id:
                print(valsn[0].message)
                break

                
code ='Console.WriteLine("Apple");'
code = '''package jv.comping;

import java.util.Dictionary;

public class comp{
    public static void main(String[] args) {
        System.out.println("Hello, World!");

        for(int i = 0; i < 10; i++){
            System.out.println("Current out" + i);
        }
    }

    public static String[][] mapper(String[][] maps1, float Apple, int Apple2, double Apple3){

        int[] val = new int[20];
        return maps1;

    }
}
'''
#lang = 'Java'
#dd = asyncio.run(runner(lang, code))