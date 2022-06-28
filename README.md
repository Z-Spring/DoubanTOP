*Last update Time: 2022-06-28 18:32:45*

**If you first use this , please delete req.Header.Add("cookie","....") in fetch.go** 

🎁 you can also read [Douban-Movie250](https://github.com/Z-Spring/Douban-Movie250) which achieves the same features but native html to parse.
```json
[
	 {
	  "id": "1",
	  "name": "肖申克的救赎 The Shawshank Redemption",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p480747492.jpg",
	  "rate": "9.7",
	  "quote": "希望让人自由。",
	  "info": "导演: 弗兰克·德拉邦特 Frank Darabont\t\t\t主演: 蒂姆·罗宾斯 Tim Robbins /...",
	  "type": "1994\t/\t美国\t/\t犯罪 剧情"
	 },
	 {
	  "id": "2",
	  "name": "霸王别姬",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561716440.jpg",
	  "rate": "9.6",
	  "quote": "风华绝代。",
	  "info": "导演: 陈凯歌 Kaige Chen\t\t\t主演: 张国荣 Leslie Cheung / 张丰毅 Fengyi Zha...",
	  "type": "1993\t/\t中国大陆 中国香港\t/\t剧情 爱情 同性"
	 },
	 {
	  "id": "3",
	  "name": "阿甘正传 Forrest Gump",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2372307693.jpg",
	  "rate": "9.5",
	  "quote": "一部美国近现代史。",
	  "info": "导演: 罗伯特·泽米吉斯 Robert Zemeckis\t\t\t主演: 汤姆·汉克斯 Tom Hanks / ...",
	  "type": "1994\t/\t美国\t/\t剧情 爱情"
	 },
	 {
	  "id": "4",
	  "name": "泰坦尼克号 Titanic",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p457760035.jpg",
	  "rate": "9.4",
	  "quote": "失去的才是永恒的。 ",
	  "info": "导演: 詹姆斯·卡梅隆 James Cameron\t\t\t主演: 莱昂纳多·迪卡普里奥 Leonardo...",
	  "type": "1997\t/\t美国 墨西哥 澳大利亚 加拿大\t/\t剧情 爱情 灾难"
	 },
	 {
	  "id": "5",
	  "name": "这个杀手不太冷 Léon",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p511118051.jpg",
	  "rate": "9.4",
	  "quote": "怪蜀黍和小萝莉不得不说的故事。",
	  "info": "导演: 吕克·贝松 Luc Besson\t\t\t主演: 让·雷诺 Jean Reno / 娜塔莉·波特曼 ...",
	  "type": "1994\t/\t法国 美国\t/\t剧情 动作 犯罪"
	 },
	 {
	  "id": "6",
	  "name": "美丽人生 La vita è bella",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2578474613.jpg",
	  "rate": "9.6",
	  "quote": "最美的谎言。",
	  "info": "导演: 罗伯托·贝尼尼 Roberto Benigni\t\t\t主演: 罗伯托·贝尼尼 Roberto Beni...",
	  "type": "1997\t/\t意大利\t/\t剧情 喜剧 爱情 战争"
	 },
	 {
	  "id": "7",
	  "name": "千与千寻 千と千尋の神隠し",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2557573348.jpg",
	  "rate": "9.4",
	  "quote": "最好的宫崎骏，最好的久石让。 ",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 柊瑠美 Rumi Hîragi / 入野自由 Miy...",
	  "type": "2001\t/\t日本\t/\t剧情 动画 奇幻"
	 },
	 {
	  "id": "8",
	  "name": "辛德勒的名单 Schindler's List",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p492406163.jpg",
	  "rate": "9.6",
	  "quote": "拯救一个人，就是拯救整个世界。",
	  "info": "导演: 史蒂文·斯皮尔伯格 Steven Spielberg\t\t\t主演: 连姆·尼森 Liam Neeson...",
	  "type": "1993\t/\t美国\t/\t剧情 历史 战争"
	 },
	 {
	  "id": "9",
	  "name": "盗梦空间 Inception",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p513344864.jpg",
	  "rate": "9.4",
	  "quote": "诺兰给了我们一场无法盗取的梦。",
	  "info": "导演: 克里斯托弗·诺兰 Christopher Nolan\t\t\t主演: 莱昂纳多·迪卡普里奥 Le...",
	  "type": "2010\t/\t美国 英国\t/\t剧情 科幻 悬疑 冒险"
	 },
	 {
	  "id": "10",
	  "name": "星际穿越 Interstellar",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2614988097.jpg",
	  "rate": "9.4",
	  "quote": "爱是一种力量，让我们超越时空感知它的存在。",
	  "info": "导演: 克里斯托弗·诺兰 Christopher Nolan\t\t\t主演: 马修·麦康纳 Matthew Mc...",
	  "type": "2014\t/\t美国 英国 加拿大\t/\t剧情 科幻 冒险"
	 },
	 {
	  "id": "11",
	  "name": "忠犬八公的故事 Hachi: A Dog's Tale",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2587099240.jpg",
	  "rate": "9.4",
	  "quote": "永远都不能忘记你所爱的人。",
	  "info": "导演: 莱塞·霍尔斯道姆 Lasse Hallström\t\t\t主演: 理查·基尔 Richard Ger...",
	  "type": "2009\t/\t美国 英国\t/\t剧情"
	 },
	 {
	  "id": "12",
	  "name": "楚门的世界 The Truman Show",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p479682972.jpg",
	  "rate": "9.3",
	  "quote": "如果再也不能见到你，祝你早安，午安，晚安。",
	  "info": "导演: 彼得·威尔 Peter Weir\t\t\t主演: 金·凯瑞 Jim Carrey / 劳拉·琳妮 Lau...",
	  "type": "1998\t/\t美国\t/\t剧情 科幻"
	 },
	 {
	  "id": "13",
	  "name": "海上钢琴师 La leggenda del pianista sull'oceano",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2574551676.jpg",
	  "rate": "9.3",
	  "quote": "每个人都要走一条自己坚定了的路，就算是粉身碎骨。 ",
	  "info": "导演: 朱塞佩·托纳多雷 Giuseppe Tornatore\t\t\t主演: 蒂姆·罗斯 Tim Roth / ...",
	  "type": "1998\t/\t意大利\t/\t剧情 音乐"
	 },
	 {
	  "id": "14",
	  "name": "三傻大闹宝莱坞 3 Idiots",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p579729551.jpg",
	  "rate": "9.2",
	  "quote": "英俊版憨豆，高情商版谢耳朵。",
	  "info": "导演: 拉库马·希拉尼 Rajkumar Hirani\t\t\t主演: 阿米尔·汗 Aamir Khan / 卡...",
	  "type": "2009\t/\t印度\t/\t剧情 喜剧 爱情 歌舞"
	 },
	 {
	  "id": "15",
	  "name": "机器人总动员 WALL·E",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1461851991.jpg",
	  "rate": "9.3",
	  "quote": "小瓦力，大人生。",
	  "info": "导演: 安德鲁·斯坦顿 Andrew Stanton\t\t\t主演: 本·贝尔特 Ben Burtt / 艾丽...",
	  "type": "2008\t/\t美国\t/\t科幻 动画 冒险"
	 },
	 {
	  "id": "16",
	  "name": "放牛班的春天 Les choristes",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1910824951.jpg",
	  "rate": "9.3",
	  "quote": "天籁一般的童声，是最接近上帝的存在。 ",
	  "info": "导演: 克里斯托夫·巴拉蒂 Christophe Barratier\t\t\t主演: 让-巴蒂斯特·莫尼...",
	  "type": "2004\t/\t法国 瑞士 德国\t/\t剧情 喜剧 音乐"
	 },
	 {
	  "id": "17",
	  "name": "无间道 無間道",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2564556863.jpg",
	  "rate": "9.3",
	  "quote": "香港电影史上永不过时的杰作。",
	  "info": "导演: 刘伟强 / 麦兆辉\t\t\t主演: 刘德华 / 梁朝伟 / 黄秋生",
	  "type": "2002\t/\t中国香港\t/\t剧情 犯罪 惊悚"
	 },
	 {
	  "id": "18",
	  "name": "疯狂动物城 Zootopia",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2614500649.jpg",
	  "rate": "9.2",
	  "quote": "迪士尼给我们营造的乌托邦就是这样，永远善良勇敢，永远出乎意料。",
	  "info": "导演: 拜伦·霍华德 Byron Howard / 瑞奇·摩尔 Rich Moore\t\t\t主演: 金妮弗·...",
	  "type": "2016\t/\t美国\t/\t喜剧 动画 冒险"
	 },
	 {
	  "id": "19",
	  "name": "大话西游之大圣娶亲 西遊記大結局之仙履奇緣",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2455050536.jpg",
	  "rate": "9.2",
	  "quote": "一生所爱。",
	  "info": "导演: 刘镇伟 Jeffrey Lau\t\t\t主演: 周星驰 Stephen Chow / 吴孟达 Man Tat Ng...",
	  "type": "1995\t/\t中国香港 中国大陆\t/\t喜剧 爱情 奇幻 古装"
	 },
	 {
	  "id": "20",
	  "name": "熔炉 도가니",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1363250216.jpg",
	  "rate": "9.3",
	  "quote": "我们一路奋战不是为了改变世界，而是为了不让世界改变我们。",
	  "info": "导演: 黄东赫 Dong-hyuk Hwang\t\t\t主演: 孔侑 Yoo Gong / 郑有美 Yu-mi Jung /...",
	  "type": "2011\t/\t韩国\t/\t剧情"
	 },
	 {
	  "id": "21",
	  "name": "控方证人 Witness for the Prosecution",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1505392928.jpg",
	  "rate": "9.6",
	  "quote": "比利·怀德满分作品。",
	  "info": "导演: 比利·怀尔德 Billy Wilder\t\t\t主演: 泰隆·鲍华 Tyrone Power / 玛琳·...",
	  "type": "1957\t/\t美国\t/\t剧情 犯罪 悬疑"
	 },
	 {
	  "id": "22",
	  "name": "教父 The Godfather",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p616779645.jpg",
	  "rate": "9.3",
	  "quote": "千万不要记恨你的对手，这样会让你失去理智。",
	  "info": "导演: 弗朗西斯·福特·科波拉 Francis Ford Coppola\t\t\t主演: 马龙·白兰度 M...",
	  "type": "1972\t/\t美国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "23",
	  "name": "当幸福来敲门 The Pursuit of Happyness",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2614359276.jpg",
	  "rate": "9.2",
	  "quote": "平民励志片。 ",
	  "info": "导演: 加布里尔·穆奇诺 Gabriele Muccino\t\t\t主演: 威尔·史密斯 Will Smith ...",
	  "type": "2006\t/\t美国\t/\t剧情 传记 家庭"
	 },
	 {
	  "id": "24",
	  "name": "触不可及 Intouchables",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1454261925.jpg",
	  "rate": "9.3",
	  "quote": "满满温情的高雅喜剧。",
	  "info": "导演: 奥利维·那卡什 Olivier Nakache / 艾力克·托兰达 Eric Toledano\t\t\t主...",
	  "type": "2011\t/\t法国\t/\t剧情 喜剧"
	 },
	 {
	  "id": "25",
	  "name": "怦然心动 Flipped",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p501177648.jpg",
	  "rate": "9.1",
	  "quote": "真正的幸福是来自内心深处。",
	  "info": "导演: 罗伯·莱纳 Rob Reiner\t\t\t主演: 玛德琳·卡罗尔 Madeline Carroll / 卡...",
	  "type": "2010\t/\t美国\t/\t剧情 喜剧 爱情"
	 },
	 {
	  "id": "26",
	  "name": "龙猫 となりのトトロ",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2540924496.jpg",
	  "rate": "9.2",
	  "quote": "人人心中都有个龙猫，童年就永远不会消失。",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 日高法子 Noriko Hidaka / 坂本千夏 Ch...",
	  "type": "1988\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "27",
	  "name": "末代皇帝 The Last Emperor",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p452089833.jpg",
	  "rate": "9.3",
	  "quote": "“不要跟我比惨，我比你更惨”再适合这部电影不过了。",
	  "info": "导演: 贝纳尔多·贝托鲁奇 Bernardo Bertolucci\t\t\t主演: 尊龙 John Lone / 陈...",
	  "type": "1987\t/\t英国 意大利 中国大陆 法国\t/\t剧情 传记 历史"
	 },
	 {
	  "id": "28",
	  "name": "寻梦环游记 Coco",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2614500706.jpg",
	  "rate": "9.1",
	  "quote": "死亡不是真的逝去，遗忘才是永恒的消亡。",
	  "info": "导演: 李·昂克里奇 Lee Unkrich / 阿德里安·莫利纳 Adrian Molina\t\t\t主演: ...",
	  "type": "2017\t/\t美国\t/\t喜剧 动画 奇幻 音乐"
	 },
	 {
	  "id": "29",
	  "name": "蝙蝠侠：黑暗骑士 The Dark Knight",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p462657443.jpg",
	  "rate": "9.2",
	  "quote": "无尽的黑暗。",
	  "info": "导演: 克里斯托弗·诺兰 Christopher Nolan\t\t\t主演: 克里斯蒂安·贝尔 Christ...",
	  "type": "2008\t/\t美国 英国\t/\t剧情 动作 科幻 犯罪 惊悚"
	 },
	 {
	  "id": "30",
	  "name": "活着",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2597919477.jpg",
	  "rate": "9.3",
	  "quote": "张艺谋最好的电影。",
	  "info": "导演: 张艺谋 Yimou Zhang\t\t\t主演: 葛优 You Ge / 巩俐 Li Gong / 姜武 Wu Jiang",
	  "type": "1994\t/\t中国大陆 中国香港\t/\t剧情 历史 家庭"
	 },
	 {
	  "id": "31",
	  "name": "哈利·波特与魔法石 Harry Potter and the Sorcerer's Stone",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2614949805.jpg",
	  "rate": "9.1",
	  "quote": "童话世界的开端。",
	  "info": "导演: Chris Columbus\t\t\t主演: Daniel Radcliffe / Emma Watson / Rupert Grint",
	  "type": "2001\t/\t美国 英国\t/\t奇幻 冒险"
	 },
	 {
	  "id": "32",
	  "name": "指环王3：王者无敌 The Lord of the Rings: The Return of the King",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2642829472.jpg",
	  "rate": "9.3",
	  "quote": "史诗的终章。",
	  "info": "导演: 彼得·杰克逊 Peter Jackson\t\t\t主演: 伊利亚·伍德 Elijah Wood / 西恩...",
	  "type": "2003\t/\t美国 新西兰\t/\t剧情 动作 奇幻 冒险"
	 },
	 {
	  "id": "33",
	  "name": "乱世佳人 Gone with the Wind",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p1963126880.jpg",
	  "rate": "9.3",
	  "quote": "Tomorrow is another day.",
	  "info": "导演: 维克多·弗莱明 Victor Fleming / 乔治·库克 George Cukor\t\t\t主演: 费...",
	  "type": "1939\t/\t美国\t/\t剧情 历史 爱情 战争"
	 },
	 {
	  "id": "34",
	  "name": "素媛 소원",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2118532944.jpg",
	  "rate": "9.3",
	  "quote": "受过伤害的人总是笑得最开心，因为他们不愿意让身边的人承受一样的痛苦。",
	  "info": "导演: 李濬益 Jun-ik Lee\t\t\t主演: 薛景求 Kyung-gu Sol / 严志媛 Ji-won Uhm ...",
	  "type": "2013\t/\t韩国\t/\t剧情"
	 },
	 {
	  "id": "35",
	  "name": "飞屋环游记 Up",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2363116942.jpg",
	  "rate": "9.1",
	  "quote": "最后那些最无聊的事情，才是最值得怀念的。 ",
	  "info": "导演: 彼特·道格特 Pete Docter / 鲍勃·彼德森 Bob Peterson\t\t\t主演: 爱德...",
	  "type": "2009\t/\t美国\t/\t剧情 喜剧 动画 冒险"
	 },
	 {
	  "id": "36",
	  "name": "摔跤吧！爸爸 Dangal",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2401676338.jpg",
	  "rate": "9.0",
	  "quote": "你不是在为你一个人战斗，你要让千千万万的女性看到女生并不是只能相夫教子。",
	  "info": "导演: 涅提·蒂瓦里 Nitesh Tiwari\t\t\t主演: 阿米尔·汗 Aamir Khan / 法缇玛...",
	  "type": "2016\t/\t印度\t/\t剧情 传记 运动 家庭"
	 },
	 {
	  "id": "37",
	  "name": "我不是药神",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2561305376.jpg",
	  "rate": "9.0",
	  "quote": "对我们国家而言，这样的电影多一部是一部。",
	  "info": "导演: 文牧野 Muye Wen\t\t\t主演: 徐峥 Zheng Xu / 王传君 Chuanjun Wang / 周...",
	  "type": "2018\t/\t中国大陆\t/\t剧情 喜剧"
	 },
	 {
	  "id": "38",
	  "name": "何以为家 كفرناحوم",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2555295759.jpg",
	  "rate": "9.1",
	  "quote": "凝视卑弱生命，用电影改变命运。",
	  "info": "导演: 娜丁·拉巴基 Nadine Labaki\t\t\t主演: 扎因·拉费阿 Zain al-Rafeea / ...",
	  "type": "2018\t/\t黎巴嫩 美国 法国 塞浦路斯 卡塔尔 英国\t/\t剧情"
	 },
	 {
	  "id": "39",
	  "name": "十二怒汉 12 Angry Men",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2173577632.jpg",
	  "rate": "9.4",
	  "quote": "1957年的理想主义。 ",
	  "info": "导演: Sidney Lumet\t\t\t主演: 亨利·方达 Henry Fonda / 马丁·鲍尔萨姆 Marti...",
	  "type": "1957\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "40",
	  "name": "哈尔的移动城堡 ハウルの動く城",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2174346180.jpg",
	  "rate": "9.1",
	  "quote": "带着心爱的人在天空飞翔。",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 倍赏千惠子 Chieko Baishô / 木村拓...",
	  "type": "2004\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "41",
	  "name": "少年派的奇幻漂流 Life of Pi",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1784592701.jpg",
	  "rate": "9.1",
	  "quote": "瑰丽壮观、无人能及的冒险之旅。",
	  "info": "导演: 李安 Ang Lee\t\t\t主演: 苏拉·沙玛 Suraj Sharma / 伊尔凡·可汗 Irrfan...",
	  "type": "2012\t/\t美国 中国台湾 英国 加拿大\t/\t剧情 奇幻 冒险"
	 },
	 {
	  "id": "42",
	  "name": "鬼子来了",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2553104888.jpg",
	  "rate": "9.3",
	  "quote": "对敌人的仁慈，就是对自己残忍。",
	  "info": "导演: 姜文 Wen Jiang\t\t\t主演: 姜文 Wen Jiang / 香川照之 Teruyuki Kagawa /...",
	  "type": "2000\t/\t中国大陆\t/\t剧情 喜剧"
	 },
	 {
	  "id": "43",
	  "name": "猫鼠游戏 Catch Me If You Can",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p453924541.jpg",
	  "rate": "9.1",
	  "quote": "骗子大师和执著警探的你追我跑故事。 ",
	  "info": "导演: 史蒂文·斯皮尔伯格 Steven Spielberg\t\t\t主演: 莱昂纳多·迪卡普里奥 L...",
	  "type": "2002\t/\t美国 加拿大\t/\t传记 犯罪 剧情"
	 },
	 {
	  "id": "44",
	  "name": "大话西游之月光宝盒 西遊記第壹佰零壹回之月光寶盒",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2561721372.jpg",
	  "rate": "9.0",
	  "quote": "旷古烁今。",
	  "info": "导演: 刘镇伟 Jeffrey Lau\t\t\t主演: 周星驰 Stephen Chow / 吴孟达 Man Tat Ng...",
	  "type": "1995\t/\t中国香港 中国大陆\t/\t喜剧 爱情 奇幻 古装"
	 },
	 {
	  "id": "45",
	  "name": "天空之城 天空の城ラピュタ",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1446261379.jpg",
	  "rate": "9.1",
	  "quote": "对天空的追逐，永不停止。 ",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 田中真弓 Mayumi Tanaka / 横泽启子 Ke...",
	  "type": "1986\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "46",
	  "name": "钢琴家 The Pianist",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p792376093.jpg",
	  "rate": "9.3",
	  "quote": "音乐能化解仇恨。",
	  "info": "导演: 罗曼·波兰斯基 Roman Polanski\t\t\t主演: 艾德里安·布洛迪 Adrien Brod...",
	  "type": "2002\t/\t英国 法国 波兰 德国\t/\t剧情 传记 战争 音乐"
	 },
	 {
	  "id": "47",
	  "name": "闻香识女人 Scent of a Woman",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2550757929.jpg",
	  "rate": "9.1",
	  "quote": "史上最美的探戈。",
	  "info": "导演: 马丁·布莱斯 Martin Brest\t\t\t主演: 阿尔·帕西诺 Al Pacino / 克里斯...",
	  "type": "1992\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "48",
	  "name": "指环王2：双塔奇兵 The Lord of the Rings: The Two Towers",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2640236255.jpg",
	  "rate": "9.2",
	  "quote": "承前启后的史诗篇章。",
	  "info": "导演: 彼得·杰克逊 Peter Jackson\t\t\t主演: 伊利亚·伍德 Elijah Wood / 西恩...",
	  "type": "2002\t/\t美国 新西兰\t/\t剧情 动作 奇幻 冒险"
	 },
	 {
	  "id": "49",
	  "name": "天堂电影院 Nuovo Cinema Paradiso",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2653054340.jpg",
	  "rate": "9.2",
	  "quote": "那些吻戏，那些青春，都在影院的黑暗里被泪水冲刷得无比清晰。",
	  "info": "导演: 朱塞佩·托纳多雷 Giuseppe Tornatore\t\t\t主演: 菲利普·努瓦雷 Philipp...",
	  "type": "1988\t/\t意大利 法国\t/\t剧情 爱情"
	 },
	 {
	  "id": "50",
	  "name": "让子弹飞",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1512562287.jpg",
	  "rate": "8.9",
	  "quote": "你给我翻译翻译，神马叫做TMD的惊喜。",
	  "info": "导演: 姜文 Wen Jiang\t\t\t主演: 姜文 Wen Jiang / 葛优 You Ge / 周润发 Yun-F...",
	  "type": "2010\t/\t中国大陆 中国香港\t/\t剧情 喜剧 动作 西部"
	 },
	 {
	  "id": "51",
	  "name": "罗马假日 Roman Holiday",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2189265085.jpg",
	  "rate": "9.1",
	  "quote": "爱情哪怕只有一天。",
	  "info": "导演: 威廉·惠勒 William Wyler\t\t\t主演: 奥黛丽·赫本 Audrey Hepburn / 格...",
	  "type": "1953\t/\t美国\t/\t喜剧 剧情 爱情"
	 },
	 {
	  "id": "52",
	  "name": "海蒂和爷爷 Heidi",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2554525534.jpg",
	  "rate": "9.3",
	  "quote": "如果生活中有什么使你感到快乐，那就去做吧！不要管别人说什么。",
	  "info": "导演: 阿兰·葛斯彭纳 Alain Gsponer\t\t\t主演: 阿努克·斯特芬 Anuk Steffen /...",
	  "type": "2015\t/\t德国 瑞士\t/\t剧情 冒险 家庭"
	 },
	 {
	  "id": "53",
	  "name": "黑客帝国 The Matrix",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p451926968.jpg",
	  "rate": "9.1",
	  "quote": "视觉革命。",
	  "info": "导演: 安迪·沃卓斯基 Andy Wachowski / 拉娜·沃卓斯基 Lana Wachowski\t\t\t主...",
	  "type": "1999\t/\t美国\t/\t动作 科幻"
	 },
	 {
	  "id": "54",
	  "name": "指环王1：护戒使者 The Lord of the Rings: The Fellowship of the Ring",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2373910134.jpg",
	  "rate": "9.1",
	  "quote": "传说的开始。",
	  "info": "导演: 彼得·杰克逊 Peter Jackson\t\t\t主演: 伊利亚·伍德 Elijah Wood / 西恩...",
	  "type": "2001\t/\t新西兰 美国\t/\t剧情 动作 奇幻 冒险"
	 },
	 {
	  "id": "55",
	  "name": "大闹天宫",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2184505167.jpg",
	  "rate": "9.4",
	  "quote": "经典之作，历久弥新。",
	  "info": "导演: 万籁鸣 Laiming Wan\t\t\t主演: 邱岳峰 Yuefeng Qiu / 富润生 Runsheng Fu...",
	  "type": "1961(中国大陆) / 1964(中国大陆) / 1978(中国大陆)\t/\t中国大陆\t/\t剧情 动画 奇幻 古装"
	 },
	 {
	  "id": "56",
	  "name": "死亡诗社 Dead Poets Society",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2575465690.jpg",
	  "rate": "9.1",
	  "quote": "当一个死水般的体制内出现一个活跃的变数时，所有的腐臭都站在了光明的对面。",
	  "info": "导演: 彼得·威尔 Peter Weir\t\t\t主演: 罗宾·威廉姆斯 Robin Williams / 罗伯...",
	  "type": "1989\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "57",
	  "name": "教父2 The Godfather: Part II",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2194138787.jpg",
	  "rate": "9.2",
	  "quote": "优雅的孤独。",
	  "info": "导演: 弗朗西斯·福特·科波拉 Francis Ford Coppola\t\t\t主演: 阿尔·帕西诺 A...",
	  "type": "1974\t/\t美国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "58",
	  "name": "辩护人 변호인",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2158166535.jpg",
	  "rate": "9.2",
	  "quote": "电影的现实意义大过电影本身。",
	  "info": "导演: 杨宇硕 Woo-seok Yang\t\t\t主演: 宋康昊 Kang-ho Song / 金英爱 Yeong-ae...",
	  "type": "2013\t/\t韩国\t/\t剧情"
	 },
	 {
	  "id": "59",
	  "name": "绿皮书 Green Book",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2549177902.jpg",
	  "rate": "8.9",
	  "quote": "去除成见，需要勇气。",
	  "info": "导演: 彼得·法雷里 Peter Farrelly\t\t\t主演: 维果·莫腾森 Viggo Mortensen /...",
	  "type": "2018\t/\t美国 中国大陆\t/\t剧情 喜剧 传记 音乐"
	 },
	 {
	  "id": "60",
	  "name": "狮子王 The Lion King",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2277799019.jpg",
	  "rate": "9.1",
	  "quote": "动物版《哈姆雷特》。",
	  "info": "导演: Roger Allers / 罗伯·明可夫 Rob Minkoff\t\t\t主演: 乔纳森·泰勒·托马...",
	  "type": "1994\t/\t美国\t/\t动画 冒险 歌舞"
	 },
	 {
	  "id": "61",
	  "name": "搏击俱乐部 Fight Club",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1910926158.jpg",
	  "rate": "9.0",
	  "quote": "邪恶与平庸蛰伏于同一个母体，在特定的时间互相对峙。",
	  "info": "导演: 大卫·芬奇 David Fincher\t\t\t主演: 爱德华·诺顿 Edward Norton / 布拉...",
	  "type": "1999\t/\t美国 德国 意大利\t/\t剧情 动作 悬疑 惊悚"
	 },
	 {
	  "id": "62",
	  "name": "饮食男女 飲食男女",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1910899751.jpg",
	  "rate": "9.2",
	  "quote": "人生不能像做菜，把所有的料都准备好了才下锅。",
	  "info": "导演: 李安 Ang Lee\t\t\t主演: 郎雄 Sihung Lung / 杨贵媚 Kuei-Mei Yang / 吴...",
	  "type": "1994\t/\t中国台湾 美国\t/\t剧情 家庭"
	 },
	 {
	  "id": "63",
	  "name": "美丽心灵 A Beautiful Mind",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p1665997400.jpg",
	  "rate": "9.1",
	  "quote": "爱是一切逻辑和原由。",
	  "info": "导演: 朗·霍华德 Ron Howard\t\t\t主演: 罗素·克劳 Russell Crowe / 艾德·哈...",
	  "type": "2001\t/\t美国\t/\t传记 剧情"
	 },
	 {
	  "id": "64",
	  "name": "本杰明·巴顿奇事 The Curious Case of Benjamin Button",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2192535722.jpg",
	  "rate": "9.0",
	  "quote": "在时间之河里感受溺水之苦。",
	  "info": "导演: 大卫·芬奇 David Fincher\t\t\t主演: 凯特·布兰切特 Cate Blanchett / ...",
	  "type": "2008\t/\t美国\t/\t剧情 爱情 奇幻"
	 },
	 {
	  "id": "65",
	  "name": "窃听风暴 Das Leben der Anderen",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1808872109.jpg",
	  "rate": "9.2",
	  "quote": "别样人生。",
	  "info": "导演: 弗洛里安·亨克尔·冯·多纳斯马尔克 Florian Henckel von Donnersmarck\t\t\u0026n...",
	  "type": "2006\t/\t德国\t/\t剧情 悬疑"
	 },
	 {
	  "id": "66",
	  "name": "穿条纹睡衣的男孩 The Boy in the Striped Pajamas",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1473670352.jpg",
	  "rate": "9.2",
	  "quote": "尽管有些不切实际的幻想，这部电影依旧是一部感人肺腑的佳作。",
	  "info": "导演: 马克·赫尔曼 Mark Herman\t\t\t主演: 阿萨·巴特菲尔德 Asa Butterfield ...",
	  "type": "2008\t/\t英国 美国\t/\t剧情 战争"
	 },
	 {
	  "id": "67",
	  "name": "情书 Love Letter",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2648230660.jpg",
	  "rate": "8.9",
	  "quote": "暗恋的极致。",
	  "info": "导演: 岩井俊二 Shunji Iwai\t\t\t主演: 中山美穗 Miho Nakayama / 丰川悦司 Ets...",
	  "type": "1995\t/\t日本\t/\t剧情 爱情"
	 },
	 {
	  "id": "68",
	  "name": "两杆大烟枪 Lock, Stock and Two Smoking Barrels",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p792443418.jpg",
	  "rate": "9.1",
	  "quote": "4个臭皮匠顶个诸葛亮，盖·里奇果然不是盖的。",
	  "info": "导演: Guy Ritchie\t\t\t主演: Jason Flemyng / Dexter Fletcher / Nick Moran",
	  "type": "1998\t/\t英国\t/\t剧情 喜剧 犯罪"
	 },
	 {
	  "id": "69",
	  "name": "西西里的美丽传说 Malèna",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2441988159.jpg",
	  "rate": "8.9",
	  "quote": "美丽无罪。",
	  "info": "导演: 朱塞佩·托纳多雷 Giuseppe Tornatore\t\t\t主演: 莫妮卡·贝鲁奇 Monica ...",
	  "type": "2000\t/\t意大利 美国\t/\t剧情 战争 情色"
	 },
	 {
	  "id": "70",
	  "name": "看不见的客人 Contratiempo",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2498971355.jpg",
	  "rate": "8.8",
	  "quote": "你以为你以为的就是你以为的。",
	  "info": "导演: 奥里奥尔·保罗 Oriol Paulo\t\t\t主演: 马里奥·卡萨斯 Mario Casas / 阿...",
	  "type": "2016\t/\t西班牙\t/\t剧情 犯罪 悬疑 惊悚"
	 },
	 {
	  "id": "71",
	  "name": "拯救大兵瑞恩 Saving Private Ryan",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1014542496.jpg",
	  "rate": "9.1",
	  "quote": "美利坚精神输出大片No1.",
	  "info": "导演: 史蒂文·斯皮尔伯格 Steven Spielberg\t\t\t主演: 汤姆·汉克斯 Tom Hanks...",
	  "type": "1998\t/\t美国\t/\t剧情 战争"
	 },
	 {
	  "id": "72",
	  "name": "音乐之声 The Sound of Music",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p453788577.jpg",
	  "rate": "9.1",
	  "quote": "用音乐化解仇恨，让歌声串起美好。",
	  "info": "导演: 罗伯特·怀斯 Robert Wise\t\t\t主演: 朱莉·安德鲁斯 Julie Andrews / 克...",
	  "type": "1965\t/\t美国\t/\t剧情 传记 爱情 歌舞"
	 },
	 {
	  "id": "73",
	  "name": "飞越疯人院 One Flew Over the Cuckoo's Nest",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p792238287.jpg",
	  "rate": "9.1",
	  "quote": "自由万岁。",
	  "info": "导演: 米洛斯·福尔曼 Miloš Forman\t\t\t主演: 杰克·尼科尔森 Jack Nichols...",
	  "type": "1975\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "74",
	  "name": "小鞋子 بچه های آسمان",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2165511465.jpg",
	  "rate": "9.2",
	  "quote": "奔跑的孩子是天使。",
	  "info": "导演: 马基德·马基迪 Majid Majidi\t\t\t主演: 默罕默德·阿米尔·纳吉 Mohamma...",
	  "type": "1997\t/\t伊朗\t/\t剧情 儿童 家庭"
	 },
	 {
	  "id": "75",
	  "name": "阿凡达 Avatar",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2634997853.jpg",
	  "rate": "8.8",
	  "quote": "绝对意义上的美轮美奂。",
	  "info": "导演: 詹姆斯·卡梅隆 James Cameron\t\t\t主演: 萨姆·沃辛顿 Sam Worthington ...",
	  "type": "2009\t/\t美国\t/\t动作 科幻 冒险"
	 },
	 {
	  "id": "76",
	  "name": "哈利·波特与死亡圣器(下) Harry Potter and the Deathly Hallows: Part 2",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p917846733.jpg",
	  "rate": "8.9",
	  "quote": "10年的完美句点。",
	  "info": "导演: 大卫·叶茨 David Yates\t\t\t主演: 丹尼尔·雷德克里夫 Daniel Radcliffe...",
	  "type": "2011\t/\t美国 英国\t/\t奇幻 冒险"
	 },
	 {
	  "id": "77",
	  "name": "沉默的羔羊 The Silence of the Lambs",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1593414327.jpg",
	  "rate": "8.9",
	  "quote": "安东尼·霍普金斯的顶级表演。",
	  "info": "导演: 乔纳森·戴米 Jonathan Demme\t\t\t主演: 朱迪·福斯特 Jodie Foster / 安...",
	  "type": "1991\t/\t美国\t/\t剧情 犯罪 惊悚"
	 },
	 {
	  "id": "78",
	  "name": "海豚湾 The Cove",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2559579779.jpg",
	  "rate": "9.3",
	  "quote": "海豚的微笑，是世界上最高明的伪装。",
	  "info": "导演: 路易·西霍尤斯 Louie Psihoyos\t\t\t主演: Richard O'Barry / 路易·西霍...",
	  "type": "2009\t/\t美国\t/\t纪录片"
	 },
	 {
	  "id": "79",
	  "name": "致命魔术 The Prestige",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p480383375.jpg",
	  "rate": "8.9",
	  "quote": "孪生蝙蝠侠大战克隆金刚狼。",
	  "info": "导演: 克里斯托弗·诺兰 Christopher Nolan\t\t\t主演: 休·杰克曼 Hugh Jackman...",
	  "type": "2006\t/\t英国 美国\t/\t剧情 悬疑 惊悚"
	 },
	 {
	  "id": "80",
	  "name": "禁闭岛 Shutter Island",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p450262388.jpg",
	  "rate": "8.8",
	  "quote": "昔日翩翩少年，今日大腹便便。",
	  "info": "导演: Martin Scorsese\t\t\t主演: 莱昂纳多·迪卡普里奥 Leonardo DiCaprio / ...",
	  "type": "2010\t/\t美国\t/\t剧情 悬疑 惊悚"
	 },
	 {
	  "id": "81",
	  "name": "布达佩斯大饭店 The Grand Budapest Hotel",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2178872593.jpg",
	  "rate": "8.9",
	  "quote": "小清新的故事里注入了大历史的情怀。",
	  "info": "导演: 韦斯·安德森 Wes Anderson\t\t\t主演: 拉尔夫·费因斯 Ralph Fiennes / ...",
	  "type": "2014\t/\t美国 德国 英国\t/\t剧情 喜剧 冒险"
	 },
	 {
	  "id": "82",
	  "name": "蝴蝶效应 The Butterfly Effect",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2209066019.jpg",
	  "rate": "8.8",
	  "quote": "人的命运被自己瞬间的抉择改变。",
	  "info": "导演: 埃里克·布雷斯 Eric Bress / J·麦基·格鲁伯 J. Mackye Gruber\t\t\t主...",
	  "type": "2004\t/\t美国 加拿大\t/\t剧情 悬疑 科幻 惊悚"
	 },
	 {
	  "id": "83",
	  "name": "美国往事 Once Upon a Time in America",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p477229647.jpg",
	  "rate": "9.2",
	  "quote": "往事如烟，无处祭奠。",
	  "info": "导演: 赛尔乔·莱翁内 Sergio Leone\t\t\t主演: 罗伯特·德尼罗 Robert De Niro ...",
	  "type": "1984\t/\t美国 意大利\t/\t犯罪 剧情"
	 },
	 {
	  "id": "84",
	  "name": "心灵捕手 Good Will Hunting",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p480965695.jpg",
	  "rate": "8.9",
	  "quote": "人生中应该拥有这样的一段豁然开朗。",
	  "info": "导演: 格斯·范·桑特 Gus Van Sant\t\t\t主演: 马特·达蒙 Matt Damon / 罗宾·...",
	  "type": "1997\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "85",
	  "name": "低俗小说 Pulp Fiction",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1910902213.jpg",
	  "rate": "8.9",
	  "quote": "故事的高级讲法。",
	  "info": "导演: 昆汀·塔伦蒂诺 Quentin Tarantino\t\t\t主演: 约翰·特拉沃尔塔 John Tra...",
	  "type": "1994\t/\t美国\t/\t剧情 喜剧 犯罪"
	 },
	 {
	  "id": "86",
	  "name": "春光乍泄 春光乍洩",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p465939041.jpg",
	  "rate": "9.0",
	  "quote": "爱情纠缠，男女一致。",
	  "info": "导演: 王家卫 Kar Wai Wong\t\t\t主演: 张国荣 Leslie Cheung / 梁朝伟 Tony Leu...",
	  "type": "1997\t/\t中国香港 日本 韩国\t/\t剧情 爱情 同性"
	 },
	 {
	  "id": "87",
	  "name": "摩登时代 Modern Times",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2263408369.jpg",
	  "rate": "9.3",
	  "quote": "大时代中的人生，小人物的悲喜。",
	  "info": "导演: 查理·卓别林 Charles Chaplin\t\t\t主演: 查理·卓别林 Charles Chaplin ...",
	  "type": "1936\t/\t美国\t/\t剧情 喜剧 爱情"
	 },
	 {
	  "id": "88",
	  "name": "七宗罪 Se7en",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2219586434.jpg",
	  "rate": "8.8",
	  "quote": "警察抓小偷，老鼠玩死猫。",
	  "info": "导演: 大卫·芬奇 David Fincher\t\t\t主演: 摩根·弗里曼 Morgan Freeman / 布...",
	  "type": "1995\t/\t美国\t/\t剧情 犯罪 悬疑 惊悚"
	 },
	 {
	  "id": "89",
	  "name": "喜剧之王 喜劇之王",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2579932167.jpg",
	  "rate": "8.8",
	  "quote": "我是一个演员。",
	  "info": "导演: 周星驰 Stephen Chow / 李力持 Lik-Chi Lee\t\t\t主演: 周星驰 Stephen Ch...",
	  "type": "1999\t/\t中国香港\t/\t喜剧 剧情 爱情"
	 },
	 {
	  "id": "90",
	  "name": "致命ID Identity",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2558364386.jpg",
	  "rate": "8.8",
	  "quote": "最不可能的那个人永远是最可能的。",
	  "info": "导演: 詹姆斯·曼高德 James Mangold\t\t\t主演: 约翰·库萨克 John Cusack / 雷...",
	  "type": "2003\t/\t美国\t/\t剧情 悬疑 惊悚"
	 },
	 {
	  "id": "91",
	  "name": "杀人回忆 살인의 추억",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p1633113220.jpg",
	  "rate": "8.9",
	  "quote": "关于连环杀人悬案的集体回忆。",
	  "info": "导演: 奉俊昊 Joon-ho Bong\t\t\t主演: 宋康昊 Kang-ho Song / 金相庆 Sang-kyun...",
	  "type": "2003\t/\t韩国\t/\t剧情 动作 犯罪 悬疑 惊悚"
	 },
	 {
	  "id": "92",
	  "name": "功夫",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2219011938.jpg",
	  "rate": "8.7",
	  "quote": "警恶惩奸，维护世界和平这个任务就交给你了，好吗？",
	  "info": "导演: 周星驰 Stephen Chow\t\t\t主演: 周星驰 Stephen Chow / 元秋 Qiu Yuen / ...",
	  "type": "2004\t/\t中国大陆 中国香港\t/\t动作 喜剧 犯罪 奇幻"
	 },
	 {
	  "id": "93",
	  "name": "超脱 Detachment",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1305562621.jpg",
	  "rate": "9.0",
	  "quote": "穷尽一生，我们要学会的，不过是彼此拥抱。",
	  "info": "导演: 托尼·凯耶 Tony Kaye\t\t\t主演: 艾德里安·布洛迪 Adrien Brody / 马西...",
	  "type": "2011\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "94",
	  "name": "哈利·波特与阿兹卡班的囚徒 Harry Potter and the Prisoner of Azkaban",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1910812549.jpg",
	  "rate": "8.9",
	  "quote": "不一样的导演，不一样的哈利·波特。",
	  "info": "导演: Alfonso Cuarón\t\t\t主演: 丹尼尔·雷德克里夫 Daniel Radcliffe / Emma...",
	  "type": "2004\t/\t英国 美国\t/\t奇幻 冒险"
	 },
	 {
	  "id": "95",
	  "name": "加勒比海盗 Pirates of the Caribbean: The Curse of the Black Pearl",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1596085504.jpg",
	  "rate": "8.8",
	  "quote": "约翰尼·德普的独角戏。",
	  "info": "导演: 戈尔·维宾斯基 Gore Verbinski\t\t\t主演: 约翰尼·德普 Johnny Depp / ...",
	  "type": "2003\t/\t美国\t/\t动作 冒险 奇幻"
	 },
	 {
	  "id": "96",
	  "name": "红辣椒 パプリカ",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p456825720.jpg",
	  "rate": "9.1",
	  "quote": "梦的勾结。",
	  "info": "导演: 今敏 Satoshi Kon\t\t\t主演: 林原惠美 Megumi Hayashibara / 江守彻 Toru...",
	  "type": "2006\t/\t日本\t/\t动画 悬疑 科幻 惊悚"
	 },
	 {
	  "id": "97",
	  "name": "被嫌弃的松子的一生 嫌われ松子の一生",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p884763596.jpg",
	  "rate": "8.9",
	  "quote": "以戏谑来戏谑戏谑。",
	  "info": "导演: 中岛哲也 Tetsuya Nakashima\t\t\t主演: 中谷美纪 Miki Nakatani / 瑛太 E...",
	  "type": "2006\t/\t日本\t/\t剧情 歌舞"
	 },
	 {
	  "id": "98",
	  "name": "狩猎 Jagten",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1546987967.jpg",
	  "rate": "9.1",
	  "quote": "人言可畏。",
	  "info": "导演: 托马斯·温特伯格 Thomas Vinterberg\t\t\t主演: 麦斯·米科尔森 Mads Mik...",
	  "type": "2012\t/\t丹麦 瑞典\t/\t剧情"
	 },
	 {
	  "id": "99",
	  "name": "请以你的名字呼唤我 Call Me by Your Name",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2505525050.jpg",
	  "rate": "8.9",
	  "quote": "沉醉在电影的情感和视听氛围中无法自拔。",
	  "info": "导演: 卢卡·瓜达尼诺 Luca Guadagnino\t\t\t主演: 艾米·汉莫 Armie Hammer / ...",
	  "type": "2017\t/\t意大利 法国 巴西 美国\t/\t剧情 爱情 同性"
	 },
	 {
	  "id": "100",
	  "name": "7号房的礼物 7번방의 선물",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1816276065.jpg",
	  "rate": "8.9",
	  "quote": "《我是山姆》的《美丽人生》。",
	  "info": "导演: 李焕庆 Hwan-kyeong Lee\t\t\t主演: 柳承龙 Seung-yong Ryoo / 朴信惠 Shi...",
	  "type": "2013\t/\t韩国\t/\t剧情 喜剧 家庭"
	 },
	 {
	  "id": "101",
	  "name": "剪刀手爱德华 Edward Scissorhands",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p480956937.jpg",
	  "rate": "8.7",
	  "quote": "浪漫忧郁的成人童话。",
	  "info": "导演: 蒂姆·波顿 Tim Burton\t\t\t主演: 约翰尼·德普 Johnny Depp / 薇诺娜·...",
	  "type": "1990\t/\t美国\t/\t剧情 奇幻 爱情"
	 },
	 {
	  "id": "102",
	  "name": "断背山 Brokeback Mountain",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2604889017.jpg",
	  "rate": "8.8",
	  "quote": "每个人心中都有一座断背山。",
	  "info": "导演: 李安 Ang Lee\t\t\t主演: 希斯·莱杰 Heath Ledger / 杰克·吉伦哈尔 Jake...",
	  "type": "2005\t/\t美国 加拿大\t/\t剧情 爱情 同性 家庭"
	 },
	 {
	  "id": "103",
	  "name": "勇敢的心 Braveheart",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2004174709.jpg",
	  "rate": "8.9",
	  "quote": "史诗大片的典范。",
	  "info": "导演: 梅尔·吉布森 Mel Gibson\t\t\t主演: 梅尔·吉布森 Mel Gibson / 苏菲·玛...",
	  "type": "1995\t/\t美国\t/\t动作 传记 剧情 历史 战争"
	 },
	 {
	  "id": "104",
	  "name": "唐伯虎点秋香 唐伯虎點秋香",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2357915564.jpg",
	  "rate": "8.7",
	  "quote": "华太师是黄霑，吴镇宇四大才子之一。",
	  "info": "导演: 李力持 Lik-Chi Lee\t\t\t主演: 周星驰 Stephen Chow / 巩俐 Li Gong / 陈...",
	  "type": "1993\t/\t中国香港\t/\t喜剧 爱情 古装"
	 },
	 {
	  "id": "105",
	  "name": "入殓师 おくりびと",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2707581855.jpg",
	  "rate": "8.9",
	  "quote": "死可能是一道门，逝去并不是终结，而是超越，走向下一程。",
	  "info": "导演: 泷田洋二郎 Yôjirô Takita\t\t\t主演: 本木雅弘 Masahiro Motoki / ...",
	  "type": "2008\t/\t日本\t/\t剧情"
	 },
	 {
	  "id": "106",
	  "name": "第六感 The Sixth Sense",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2220184425.jpg",
	  "rate": "8.9",
	  "quote": "深入内心的恐怖，出人意料的结局。",
	  "info": "导演: M·奈特·沙马兰 M. Night Shyamalan\t\t\t主演: 布鲁斯·威利斯 Bruce Wi...",
	  "type": "1999\t/\t美国\t/\t剧情 悬疑 惊悚"
	 },
	 {
	  "id": "107",
	  "name": "哈利·波特与密室 Harry Potter and the Chamber of Secrets",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p1082651990.jpg",
	  "rate": "8.8",
	  "quote": "魔法的密室之门已打开...",
	  "info": "导演: Chris Columbus\t\t\t主演: 丹尼尔·雷德克里夫 Daniel Radcliffe / 艾玛...",
	  "type": "2002\t/\t英国 美国\t/\t奇幻 冒险"
	 },
	 {
	  "id": "108",
	  "name": "天使爱美丽 Le fabuleux destin d'Amélie Poulain",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2447590313.jpg",
	  "rate": "8.7",
	  "quote": "法式小清新。 ",
	  "info": "导演: 让-皮埃尔·热内 Jean-Pierre Jeunet\t\t\t主演: 奥黛丽·塔图 Audrey Tau...",
	  "type": "2001\t/\t法国 德国\t/\t剧情 喜剧 爱情"
	 },
	 {
	  "id": "109",
	  "name": "一一",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2567845803.jpg",
	  "rate": "9.1",
	  "quote": "我们都曾经是一一。",
	  "info": "导演: 杨德昌 Edward Yang\t\t\t主演: 吴念真 / 李凯莉 Kelly Lee / 金燕玲 Elai...",
	  "type": "2000\t/\t中国台湾 日本\t/\t剧情 爱情 家庭"
	 },
	 {
	  "id": "110",
	  "name": "爱在黎明破晓前 Before Sunrise",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2555762374.jpg",
	  "rate": "8.8",
	  "quote": "缘分是个连绵词，最美不过一瞬。",
	  "info": "导演: 理查德·林克莱特 Richard Linklater\t\t\t主演: 伊桑·霍克 Ethan Hawke ...",
	  "type": "1995\t/\t美国 奥地利 瑞士\t/\t剧情 爱情"
	 },
	 {
	  "id": "111",
	  "name": "重庆森林 重慶森林",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p792381411.jpg",
	  "rate": "8.8",
	  "quote": "寂寞没有期限。",
	  "info": "导演: 王家卫 Kar Wai Wong\t\t\t主演: 林青霞 Brigitte Lin / 金城武 Takeshi K...",
	  "type": "1994\t/\t中国香港\t/\t剧情 爱情"
	 },
	 {
	  "id": "112",
	  "name": "幽灵公主 もののけ姫",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1613191025.jpg",
	  "rate": "8.9",
	  "quote": "人与自然的战争史诗。",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 松田洋治 Yôji Matsuda / 石田百合...",
	  "type": "1997\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "113",
	  "name": "蝙蝠侠：黑暗骑士崛起 The Dark Knight Rises",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1706428744.jpg",
	  "rate": "8.8",
	  "quote": "诺兰就是保证。",
	  "info": "导演: 克里斯托弗·诺兰 Christopher Nolan\t\t\t主演: 克里斯蒂安·贝尔 Christ...",
	  "type": "2012\t/\t美国 英国\t/\t剧情 动作 科幻 犯罪 惊悚"
	 },
	 {
	  "id": "114",
	  "name": "小森林 夏秋篇 リトル・フォレスト 夏・秋",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2564498893.jpg",
	  "rate": "9.0",
	  "quote": "那些静得只能听见呼吸的日子里，你明白孤独即生活。",
	  "info": "导演: 森淳一 Junichi Mori\t\t\t主演: 桥本爱 Ai Hashimoto / 三浦贵大 Takahir...",
	  "type": "2014\t/\t日本\t/\t剧情"
	 },
	 {
	  "id": "115",
	  "name": "阳光灿烂的日子",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2564685215.jpg",
	  "rate": "8.8",
	  "quote": "一场华丽的意淫。",
	  "info": "导演: 姜文 Wen Jiang\t\t\t主演: 夏雨 Yu Xia / 宁静 Jing Ning / 陶虹 Hong Tao",
	  "type": "1994\t/\t中国大陆 中国香港\t/\t剧情 爱情"
	 },
	 {
	  "id": "116",
	  "name": "菊次郎的夏天 菊次郎の夏",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2620392435.jpg",
	  "rate": "8.9",
	  "quote": "从没见过那么流氓的温柔，从没见过那么温柔的流氓。",
	  "info": "导演: 北野武 Takeshi Kitano\t\t\t主演: 北野武 Takeshi Kitano / 关口雄介 Yus...",
	  "type": "1999\t/\t日本\t/\t剧情 喜剧"
	 },
	 {
	  "id": "117",
	  "name": "超能陆战队 Big Hero 6",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2614500883.jpg",
	  "rate": "8.7",
	  "quote": "Balalala~~~",
	  "info": "导演: 唐·霍尔 Don Hall / 克里斯·威廉姆斯 Chris Williams\t\t\t主演: 斯科特...",
	  "type": "2014\t/\t美国\t/\t喜剧 动作 科幻 动画 冒险"
	 },
	 {
	  "id": "118",
	  "name": "完美的世界 A Perfect World",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2190556408.jpg",
	  "rate": "9.1",
	  "quote": "坏人的好总是比好人的好来得更感人。",
	  "info": "导演: 克林特·伊斯特伍德 Clint Eastwood\t\t\t主演: 凯文·科斯特纳 Kevin Cos...",
	  "type": "1993\t/\t美国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "119",
	  "name": "无人知晓 誰も知らない",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p661160053.jpg",
	  "rate": "9.1",
	  "quote": "我的平常生活就是他人的幸福。",
	  "info": "导演: 是枝裕和 Hirokazu Koreeda\t\t\t主演: 柳乐优弥 Yûya Yagira / 北浦爱...",
	  "type": "2004\t/\t日本\t/\t剧情"
	 },
	 {
	  "id": "120",
	  "name": "爱在日落黄昏时 Before Sunset",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2561542458.jpg",
	  "rate": "8.9",
	  "quote": "九年后的重逢是世俗和责任的交叠，没了悸动和青涩，沧桑而温暖。",
	  "info": "导演: 理查德·林克莱特 Richard Linklater\t\t\t主演: 伊桑·霍克 Ethan Hawke ...",
	  "type": "2004\t/\t美国 法国\t/\t剧情 爱情"
	 },
	 {
	  "id": "121",
	  "name": "消失的爱人 Gone Girl",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2221768894.jpg",
	  "rate": "8.7",
	  "quote": "年度最佳date movie。",
	  "info": "导演: 大卫·芬奇 David Fincher\t\t\t主演: 本·阿弗莱克 Ben Affleck / 罗莎蒙...",
	  "type": "2014\t/\t美国\t/\t剧情 犯罪 悬疑 惊悚"
	 },
	 {
	  "id": "122",
	  "name": "借东西的小人阿莉埃蒂 借りぐらしのアリエッティ",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p617533616.jpg",
	  "rate": "8.9",
	  "quote": "曾经的那段美好会沉淀为一辈子的记忆。",
	  "info": "导演: 米林宏昌 Hiromasa Yonebayashi\t\t\t主演: 志田未来 Mirai Shida / 神木...",
	  "type": "2010\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "123",
	  "name": "小森林 冬春篇 リトル・フォレスト 冬・春",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2215147728.jpg",
	  "rate": "9.0",
	  "quote": "尊敬他人，尊敬你生活的这片土地，明白孤独是人生的常态。",
	  "info": "导演: 森淳一 Junichi Mori\t\t\t主演: 桥本爱 Ai Hashimoto / 三浦贵大 Takahir...",
	  "type": "2015\t/\t日本\t/\t剧情"
	 },
	 {
	  "id": "124",
	  "name": "甜蜜蜜",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2223011274.jpg",
	  "rate": "8.9",
	  "quote": "相逢只要一瞬间，等待却像是一辈子。",
	  "info": "导演: 陈可辛 Peter Chan\t\t\t主演: 黎明 Leon Lai / 张曼玉 Maggie Cheung / ...",
	  "type": "1996\t/\t中国香港\t/\t剧情 爱情"
	 },
	 {
	  "id": "125",
	  "name": "倩女幽魂",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2414157745.jpg",
	  "rate": "8.8",
	  "quote": "两张绝世的脸。 ",
	  "info": "导演: 程小东 Siu-Tung Ching\t\t\t主演: 张国荣 Leslie Cheung / 王祖贤 Joey W...",
	  "type": "1987\t/\t中国香港\t/\t爱情 奇幻 武侠 古装"
	 },
	 {
	  "id": "126",
	  "name": "侧耳倾听 耳をすませば",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p456692072.jpg",
	  "rate": "8.9",
	  "quote": "少女情怀总是诗。",
	  "info": "导演: 近藤喜文 Yoshifumi Kondo\t\t\t主演: 本名阳子 Youko Honna / 小林桂树 K...",
	  "type": "1995\t/\t日本\t/\t剧情 爱情 动画"
	 },
	 {
	  "id": "127",
	  "name": "幸福终点站 The Terminal",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p854757687.jpg",
	  "rate": "8.8",
	  "quote": "有时候幸福需要等一等。 ",
	  "info": "导演: 史蒂文·斯皮尔伯格 Steven Spielberg\t\t\t主演: 汤姆·汉克斯 Tom Hanks...",
	  "type": "2004\t/\t美国\t/\t喜剧 剧情 爱情"
	 },
	 {
	  "id": "128",
	  "name": "时空恋旅人 About Time",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2070153774.jpg",
	  "rate": "8.8",
	  "quote": "把每天当作最后一天般珍惜度过，积极拥抱生活，就是幸福。",
	  "info": "导演: 理查德·柯蒂斯 Richard Curtis\t\t\t主演: 多姆纳尔·格里森 Domhnall Gl...",
	  "type": "2013\t/\t英国\t/\t喜剧 爱情 奇幻"
	 },
	 {
	  "id": "129",
	  "name": "驯龙高手 How to Train Your Dragon",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2210954024.jpg",
	  "rate": "8.7",
	  "quote": "和谐的生活离不开摸头与被摸头。",
	  "info": "导演: 迪恩·德布洛斯 Dean DeBlois / 克里斯·桑德斯 Chris Sanders\t\t\t主演:...",
	  "type": "2010\t/\t美国\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "130",
	  "name": "萤火之森 蛍火の杜へ",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1675053073.jpg",
	  "rate": "8.9",
	  "quote": "触不到的恋人。",
	  "info": "导演: 大森贵弘 Takahiro Omori\t\t\t主演: 佐仓绫音 Ayane Sakura / 内山昂辉 K...",
	  "type": "2011\t/\t日本\t/\t剧情 爱情 动画 奇幻"
	 },
	 {
	  "id": "131",
	  "name": "玛丽和马克思 Mary and Max",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p488255145.jpg",
	  "rate": "8.9",
	  "quote": "你是我最好的朋友，你是我唯一的朋友 。",
	  "info": "导演: 亚当·艾略特 Adam Elliot\t\t\t主演: 托妮·科莱特 Toni Collette / 菲利...",
	  "type": "2009\t/\t澳大利亚 美国\t/\t剧情 喜剧 动画"
	 },
	 {
	  "id": "132",
	  "name": "怪兽电力公司 Monsters, Inc.",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2513247938.jpg",
	  "rate": "8.8",
	  "quote": "不要给它起名字，起了名字就有感情了。",
	  "info": "导演: 彼特·道格特 Pete Docter / 大卫·斯沃曼 David Silverman\t\t\t主演: 约...",
	  "type": "2001\t/\t美国\t/\t儿童 喜剧 动画 奇幻 冒险"
	 },
	 {
	  "id": "133",
	  "name": "教父3 The Godfather: Part III",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2169664351.jpg",
	  "rate": "9.0",
	  "quote": "任何信念的力量，都无法改变命运。",
	  "info": "导演: 弗朗西斯·福特·科波拉 Francis Ford Coppola\t\t\t主演: 阿尔·帕西诺 A...",
	  "type": "1990\t/\t美国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "134",
	  "name": "一个叫欧维的男人决定去死 En man som heter Ove",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2406624993.jpg",
	  "rate": "8.9",
	  "quote": "惠及一生的美丽。",
	  "info": "导演: 汉内斯·赫尔姆 Hannes Holm\t\t\t主演: 罗夫·拉斯加德 Rolf Lassgård...",
	  "type": "2015\t/\t瑞典\t/\t剧情"
	 },
	 {
	  "id": "135",
	  "name": "大鱼 Big Fish",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p692813374.jpg",
	  "rate": "8.8",
	  "quote": "抱着梦想而活着的人是幸福的，怀抱梦想而死去的人是不朽的。",
	  "info": "导演: 蒂姆·波顿 Tim Burton\t\t\t主演: 伊万·麦克格雷格 Ewan McGregor / 阿...",
	  "type": "2003\t/\t美国\t/\t剧情 家庭 奇幻 冒险"
	 },
	 {
	  "id": "136",
	  "name": "玩具总动员3 Toy Story 3",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1283675359.jpg",
	  "rate": "8.9",
	  "quote": "跨度十五年的欢乐与泪水。",
	  "info": "导演: 李·昂克里奇 Lee Unkrich\t\t\t主演: 汤姆·汉克斯 Tom Hanks / 蒂姆·艾...",
	  "type": "2010\t/\t美国\t/\t喜剧 动画 奇幻 冒险"
	 },
	 {
	  "id": "137",
	  "name": "寄生虫 기생충",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561439800.jpg",
	  "rate": "8.8",
	  "info": "导演: 奉俊昊 Joon-ho Bong\t\t\t主演: 宋康昊 Kang-ho Song / 李善均 Seon-gyun...",
	  "type": "2019\t/\t韩国\t/\t剧情"
	 },
	 {
	  "id": "138",
	  "name": "傲慢与偏见 Pride \u0026 Prejudice",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2016401659.jpg",
	  "rate": "8.7",
	  "quote": "爱是摈弃傲慢与偏见之后的曙光。",
	  "info": "导演: 乔·怀特 Joe Wright\t\t\t主演: 凯拉·奈特莉 Keira Knightley / 马修·...",
	  "type": "2005\t/\t法国 英国 美国\t/\t剧情 爱情"
	 },
	 {
	  "id": "139",
	  "name": "告白",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p689520756.jpg",
	  "rate": "8.8",
	  "quote": "没有一人完全善，也没有一人完全恶。",
	  "info": "导演: 中岛哲也 Tetsuya Nakashima\t\t\t主演: 松隆子 Takako Matsu / 冈田将生 ...",
	  "type": "2010\t/\t日本\t/\t剧情 悬疑"
	 },
	 {
	  "id": "140",
	  "name": "神偷奶爸 Despicable Me",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p792776858.jpg",
	  "rate": "8.6",
	  "quote": "Mr. I Don't Care其实也有Care的时候。",
	  "info": "导演: 皮艾尔·柯芬 Pierre Coffin / 克里斯·雷纳德 Chris Renaud\t\t\t主演: ...",
	  "type": "2010\t/\t美国 法国\t/\t喜剧 动画 冒险"
	 },
	 {
	  "id": "141",
	  "name": "釜山行 부산행",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2360940399.jpg",
	  "rate": "8.6",
	  "quote": "揭露人性的丧尸题材力作。",
	  "info": "导演: 延尚昊 Sang-ho Yeon\t\t\t主演: 孔侑 Yoo Gong / 郑有美 Yu-mi Jung / 马...",
	  "type": "2016\t/\t韩国\t/\t动作 惊悚 灾难"
	 },
	 {
	  "id": "142",
	  "name": "未麻的部屋 Perfect Blue",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1351050722.jpg",
	  "rate": "9.0",
	  "quote": "好的剧本是，就算你猜到了结局也猜不到全部。",
	  "info": "导演: 今敏 Satoshi Kon\t\t\t主演: 岩男润子 Junko Iwao / 松本梨香 Rica Matsu...",
	  "type": "1997\t/\t日本\t/\t动画 奇幻 惊悚"
	 },
	 {
	  "id": "143",
	  "name": "阳光姐妹淘 써니",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1374786017.jpg",
	  "rate": "8.8",
	  "quote": "再多各自牛逼的时光，也比不上一起傻逼的岁月。 ",
	  "info": "导演: 姜炯哲 Hyeong-Cheol Kang\t\t\t主演: 沈恩京 Eun-kyung Shim / 闵孝琳 Hy...",
	  "type": "2011\t/\t韩国\t/\t剧情 喜剧"
	 },
	 {
	  "id": "144",
	  "name": "射雕英雄传之东成西就 射鵰英雄傳之東成西就",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2609063925.jpg",
	  "rate": "8.7",
	  "quote": "百看不厌。 ",
	  "info": "导演: 刘镇伟 Jeffrey Lau\t\t\t主演: 梁朝伟 Tony Leung Chiu Wai / 林青霞 Bri...",
	  "type": "1993\t/\t中国香港\t/\t喜剧 奇幻 武侠 古装"
	 },
	 {
	  "id": "145",
	  "name": "被解救的姜戈 Django Unchained",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1800813767.jpg",
	  "rate": "8.8",
	  "quote": "热血沸腾，那个低俗、性感的无耻混蛋又来了。",
	  "info": "导演: 昆汀·塔伦蒂诺 Quentin Tarantino\t\t\t主演: 杰米·福克斯 Jamie Foxx /...",
	  "type": "2012\t/\t美国\t/\t剧情 动作 西部 冒险"
	 },
	 {
	  "id": "146",
	  "name": "恐怖直播 더 테러 라이브",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2054744538.jpg",
	  "rate": "8.7",
	  "quote": "恐怖分子的“秋菊打官司”。",
	  "info": "导演: 金秉祐 Byeong-woo Kim\t\t\t主演: 河正宇 Jung-woo Ha / 李璟荣 Kyeong-y...",
	  "type": "2013\t/\t韩国\t/\t剧情 犯罪 悬疑"
	 },
	 {
	  "id": "147",
	  "name": "哪吒闹海",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2532803206.jpg",
	  "rate": "9.1",
	  "quote": "想你时你在闹海。",
	  "info": "导演: 王树忱 Shuchen Wang / 严定宪 Dingxian Yan\t\t\t主演: 梁正晖 Zhenghui ...",
	  "type": "1979\t/\t中国大陆\t/\t冒险 动画 奇幻"
	 },
	 {
	  "id": "148",
	  "name": "我是山姆 I Am Sam",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p652417775.jpg",
	  "rate": "9.0",
	  "quote": "爱并不需要智商 。",
	  "info": "导演: 杰茜·尼尔森 Jessie Nelson\t\t\t主演: Sean Penn / Dakota Fanning / Mi...",
	  "type": "2001\t/\t美国\t/\t剧情 家庭"
	 },
	 {
	  "id": "149",
	  "name": "哈利·波特与火焰杯 Harry Potter and the Goblet of Fire",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p735391084.jpg",
	  "rate": "8.7",
	  "info": "导演: 迈克·内威尔 Mike Newell\t\t\t主演: 丹尼尔·雷德克里夫 Daniel Radclif...",
	  "type": "2005\t/\t英国 美国\t/\t悬疑 奇幻 冒险"
	 },
	 {
	  "id": "150",
	  "name": "血战钢锯岭 Hacksaw Ridge",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2398141939.jpg",
	  "rate": "8.7",
	  "quote": "优秀的战争片不会美化战场，不会粉饰死亡，不会矮化敌人，不会无视常识，最重要的，不会宣扬战争。",
	  "info": "导演: 梅尔·吉布森 Mel Gibson\t\t\t主演: 安德鲁·加菲尔德 Andrew Garfield /...",
	  "type": "2016\t/\t澳大利亚 美国\t/\t剧情 传记 历史 战争"
	 },
	 {
	  "id": "151",
	  "name": "头号玩家 Ready Player One",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2516578307.jpg",
	  "rate": "8.7",
	  "quote": "写给影迷，动漫迷和游戏迷的一封情书。",
	  "info": "导演: 史蒂文·斯皮尔伯格 Steven Spielberg\t\t\t主演: 泰伊·谢里丹 Tye Sheri...",
	  "type": "2018\t/\t美国\t/\t动作 科幻 冒险"
	 },
	 {
	  "id": "152",
	  "name": "新世界 신세계",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1903379979.jpg",
	  "rate": "8.9",
	  "quote": "要做就做得狠一点，这样才能活下去。",
	  "info": "导演: 朴勋政 Hoon-jung Park\t\t\t主演: 李政宰 Jung-Jae Lee / 崔岷植 Min-sik...",
	  "type": "2013\t/\t韩国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "153",
	  "name": "模仿游戏 The Imitation Game",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2255040492.jpg",
	  "rate": "8.7",
	  "quote": "他给机器起名“克里斯托弗”，因为这是他初恋的名字。",
	  "info": "导演: 莫滕·泰杜姆 Morten Tyldum\t\t\t主演: 本尼迪克特·康伯巴奇 Benedict C...",
	  "type": "2014\t/\t英国 美国\t/\t剧情 传记 战争 同性"
	 },
	 {
	  "id": "154",
	  "name": "喜宴 囍宴",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2249048907.jpg",
	  "rate": "9.0",
	  "quote": "中国家庭的喜怒哀乐忍。",
	  "info": "导演: 李安 Ang Lee\t\t\t主演: 赵文瑄 Winston Chao / 郎雄 Sihung Lung / 归亚...",
	  "type": "1993\t/\t中国台湾 美国\t/\t剧情 喜剧 爱情 同性 家庭"
	 },
	 {
	  "id": "155",
	  "name": "七武士 七人の侍",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2215886505.jpg",
	  "rate": "9.3",
	  "quote": "时代悲歌。",
	  "info": "导演: 黑泽明 Akira Kurosawa\t\t\t主演: 三船敏郎 Toshirô Mifune / 志村乔 ...",
	  "type": "1954\t/\t日本\t/\t动作 冒险 剧情"
	 },
	 {
	  "id": "156",
	  "name": "花样年华 花樣年華",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1910828286.jpg",
	  "rate": "8.7",
	  "quote": "偷情本没有这样美。",
	  "info": "导演: 王家卫 Kar Wai Wong\t\t\t主演: 梁朝伟 Tony Leung Chiu Wai / 张曼玉 Ma...",
	  "type": "2000\t/\t中国香港\t/\t剧情 爱情"
	 },
	 {
	  "id": "157",
	  "name": "黑客帝国3：矩阵革命 The Matrix Revolutions",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p443461818.jpg",
	  "rate": "8.8",
	  "quote": "不得不说，《黑客帝国》系列是商业片与科幻、哲学完美结合的典范。",
	  "info": "导演: 拉娜·沃卓斯基 Lana Wachowski / 莉莉·沃卓斯基 Lilly Wachowski\t\t\t...",
	  "type": "2003\t/\t美国\t/\t动作 科幻"
	 },
	 {
	  "id": "158",
	  "name": "头脑特工队 Inside Out",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2266293606.jpg",
	  "rate": "8.8",
	  "quote": "愿我们都不用长大，每一座城堡都能永远存在。",
	  "info": "导演: 彼特·道格特 Pete Docter / 罗纳尔多·德尔·卡门 Ronaldo Del Carmen\t\t\u0026nb...",
	  "type": "2015\t/\t美国\t/\t喜剧 动画 冒险"
	 },
	 {
	  "id": "159",
	  "name": "电锯惊魂 Saw",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p726839485.jpg",
	  "rate": "8.7",
	  "quote": "真相就在眼前。",
	  "info": "导演: 詹姆斯·温 James Wan\t\t\t主演: 雷·沃纳尔 Leigh Whannell / 加利·艾...",
	  "type": "2004\t/\t美国\t/\t悬疑 惊悚 恐怖"
	 },
	 {
	  "id": "160",
	  "name": "三块广告牌 Three Billboards Outside Ebbing, Missouri",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2510081688.jpg",
	  "rate": "8.7",
	  "quote": "怼天怼地，你走后，她与世界为敌。",
	  "info": "导演: 马丁·麦克唐纳 Martin McDonagh\t\t\t主演: 弗兰西斯·麦克多蒙德 France...",
	  "type": "2017\t/\t美国 英国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "161",
	  "name": "你的名字。 君の名は。",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2395733377.jpg",
	  "rate": "8.5",
	  "quote": "穿越错位的时空，仰望陨落的星辰，你没留下你的名字，我却无法忘记那句“我爱你”。",
	  "info": "导演: 新海诚 Makoto Shinkai\t\t\t主演: 神木隆之介 Ryûnosuke Kamiki / 上...",
	  "type": "2016\t/\t日本\t/\t剧情 爱情 动画"
	 },
	 {
	  "id": "162",
	  "name": "卢旺达饭店 Hotel Rwanda",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p470419493.jpg",
	  "rate": "8.9",
	  "quote": "当这个世界闭上双眼，他却敞开了怀抱。",
	  "info": "导演: 特瑞·乔治 Terry George\t\t\t主演: 唐·钱德尔 Don Cheadle / 苏菲·奥...",
	  "type": "2004\t/\t英国 南非 意大利 美国\t/\t剧情 传记 历史 战争"
	 },
	 {
	  "id": "163",
	  "name": "达拉斯买家俱乐部 Dallas Buyers Club",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2166160837.jpg",
	  "rate": "8.8",
	  "quote": "Jared Leto的腿比女人还美！",
	  "info": "导演: 让-马克·瓦雷 Jean-Marc Vallée\t\t\t主演: 马修·麦康纳 Matthew McCon...",
	  "type": "2013\t/\t美国\t/\t剧情 传记 同性"
	 },
	 {
	  "id": "164",
	  "name": "疯狂原始人 The Croods",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1867084027.jpg",
	  "rate": "8.7",
	  "quote": "老少皆宜，这就是好莱坞动画的魅力。",
	  "info": "导演: 科克·德·米科 Kirk De Micco / 克里斯·桑德斯 Chris Sanders\t\t\t主演...",
	  "type": "2013\t/\t美国\t/\t喜剧 动画 冒险"
	 },
	 {
	  "id": "165",
	  "name": "上帝之城 Cidade de Deus",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p455677490.jpg",
	  "rate": "9.0",
	  "quote": "被上帝抛弃了的上帝之城。",
	  "info": "导演: 费尔南多·梅里尔斯 Fernando Meirelles / 卡迪亚·兰德 Kátia Lund\t\t\t...",
	  "type": "2002\t/\t巴西 法国\t/\t犯罪 剧情"
	 },
	 {
	  "id": "166",
	  "name": "谍影重重3 The Bourne Ultimatum",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p792223507.jpg",
	  "rate": "8.8",
	  "quote": "像吃了苏打饼一样干脆的电影。",
	  "info": "导演: 保罗·格林格拉斯 Paul Greengrass\t\t\t主演: 马特·达蒙 Matt Damon / ...",
	  "type": "2007\t/\t美国 德国\t/\t动作 悬疑 惊悚"
	 },
	 {
	  "id": "167",
	  "name": "九品芝麻官",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p648370300.jpg",
	  "rate": "8.7",
	  "info": "导演: 王晶 Jing Wong\t\t\t主演: 周星驰 Stephen Chow / 吴孟达 Man Tat Ng / ...",
	  "type": "1994\t/\t中国香港 中国大陆\t/\t剧情 喜剧 古装"
	 },
	 {
	  "id": "168",
	  "name": "英雄本色",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2504997087.jpg",
	  "rate": "8.6",
	  "quote": "英雄泪短，兄弟情长。 ",
	  "info": "导演: 吴宇森 John Woo\t\t\t主演: 周润发 Yun-Fat Chow / 狄龙 Lung Ti / 张国...",
	  "type": "1986\t/\t中国香港\t/\t剧情 动作 犯罪"
	 },
	 {
	  "id": "169",
	  "name": "风之谷 風の谷のナウシカ",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1917567652.jpg",
	  "rate": "8.9",
	  "quote": "动画片的圣经。",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 岛本须美 Sumi Shimamoto / 松田洋治 Y...",
	  "type": "1984\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "170",
	  "name": "惊魂记 Psycho",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1021883305.jpg",
	  "rate": "9.0",
	  "quote": "故事的反转与反转，分裂电影的始祖。",
	  "info": "导演: 阿尔弗雷德·希区柯克 Alfred Hitchcock\t\t\t主演: 安东尼·博金斯 Antho...",
	  "type": "1960\t/\t美国\t/\t悬疑 惊悚 恐怖"
	 },
	 {
	  "id": "171",
	  "name": "心迷宫",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2276780256.jpg",
	  "rate": "8.7",
	  "quote": "荒诞讽刺，千奇百巧，抽丝剥茧，百转千回。",
	  "info": "导演: 忻钰坤 Yukun Xin\t\t\t主演: 霍卫民 Weimin Huo / 王笑天 Xiaotian Wang ...",
	  "type": "2014\t/\t中国大陆\t/\t剧情 犯罪 悬疑"
	 },
	 {
	  "id": "172",
	  "name": "纵横四海 緃横四海",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2597918718.jpg",
	  "rate": "8.8",
	  "quote": "香港浪漫主义警匪动作片的巅峰之作。",
	  "info": "导演: 吴宇森 John Woo\t\t\t主演: 周润发 Yun-Fat Chow / 张国荣 Leslie Cheung...",
	  "type": "1991\t/\t中国香港\t/\t剧情 喜剧 动作 犯罪"
	 },
	 {
	  "id": "173",
	  "name": "海街日记 海街diary",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2603364117.jpg",
	  "rate": "8.8",
	  "quote": "是枝裕和的家庭习作。",
	  "info": "导演: 是枝裕和 Hirokazu Koreeda\t\t\t主演: 绫濑遥 Haruka Ayase / 长泽雅美 M...",
	  "type": "2015\t/\t日本\t/\t剧情 家庭"
	 },
	 {
	  "id": "174",
	  "name": "岁月神偷 歲月神偷",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p456666151.jpg",
	  "rate": "8.7",
	  "quote": "岁月流逝，来日可追。",
	  "info": "导演: 罗启锐 Alex Law\t\t\t主演: 吴君如 Sandra Ng / 任达华 Simon Yam / 钟绍...",
	  "type": "2010\t/\t中国香港 中国大陆\t/\t剧情 家庭"
	 },
	 {
	  "id": "175",
	  "name": "记忆碎片 Memento",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p641688453.jpg",
	  "rate": "8.7",
	  "quote": "一个针管引发的血案。",
	  "info": "导演: 克里斯托弗·诺兰 Christopher Nolan\t\t\t主演: 盖·皮尔斯 Guy Pearce /...",
	  "type": "2000\t/\t美国\t/\t犯罪 剧情 悬疑 惊悚"
	 },
	 {
	  "id": "176",
	  "name": "忠犬八公物语 ハチ公物語",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2603716224.jpg",
	  "rate": "9.2",
	  "quote": "养狗三日，便会对你终其一生。",
	  "info": "导演: Seijirô Kôyama\t\t\t主演: 山本圭 Kei Yamamoto / 井川比佐志 Hisa...",
	  "type": "1987\t/\t日本\t/\t剧情"
	 },
	 {
	  "id": "177",
	  "name": "荒蛮故事 Relatos salvajes",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2584519452.jpg",
	  "rate": "8.8",
	  "quote": "始于荒诞，止于更荒诞。",
	  "info": "导演: 达米安·斯兹弗隆 Damián Szifron\t\t\t主演: 达里奥·葛兰帝内提 Darío...",
	  "type": "2014\t/\t阿根廷 西班牙\t/\t剧情 喜剧 犯罪"
	 },
	 {
	  "id": "178",
	  "name": "爱在午夜降临前 Before Midnight",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2074715729.jpg",
	  "rate": "8.9",
	  "quote": "所谓爱情，就是话唠一路，都不会心生腻烦，彼此嫌弃。",
	  "info": "导演: 理查德·林克莱特 Richard Linklater\t\t\t主演: 伊桑·霍克 Ethan Hawke ...",
	  "type": "2013\t/\t美国 希腊\t/\t剧情 爱情"
	 },
	 {
	  "id": "179",
	  "name": "绿里奇迹 The Green Mile",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p767586451.jpg",
	  "rate": "8.9",
	  "quote": "天使暂时离开。",
	  "info": "导演: Frank Darabont\t\t\t主演: 汤姆·汉克斯 Tom Hanks / 大卫·摩斯 David M...",
	  "type": "1999\t/\t美国\t/\t犯罪 剧情 奇幻 悬疑"
	 },
	 {
	  "id": "180",
	  "name": "色，戒",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p453716305.jpg",
	  "rate": "8.6",
	  "quote": "假戏真情，爱欲深海",
	  "info": "导演: 李安 Ang Lee\t\t\t主演: 梁朝伟 Tony Leung Chiu Wai / 汤唯 Wei Tang / ...",
	  "type": "2007\t/\t中国台湾 中国大陆 美国 中国香港\t/\t剧情 爱情 情色"
	 },
	 {
	  "id": "181",
	  "name": "爆裂鼓手 Whiplash",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2220776342.jpg",
	  "rate": "8.7",
	  "quote": "这个世界从不善待努力的人，努力了也不一定会成功，但是知道自己在努力，就是活下去的动力。",
	  "info": "导演: 达米恩·查泽雷 Damien Chazelle\t\t\t主演: 迈尔斯·特勒 Miles Teller /...",
	  "type": "2014\t/\t美国\t/\t剧情 音乐"
	 },
	 {
	  "id": "182",
	  "name": "小偷家族 万引き家族",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2530599636.jpg",
	  "rate": "8.7",
	  "quote": "我们组成了家。",
	  "info": "导演: 是枝裕和 Hirokazu Koreeda\t\t\t主演: 中川雅也 Lily Franky / 安藤樱 Sa...",
	  "type": "2018\t/\t日本\t/\t剧情 犯罪 家庭"
	 },
	 {
	  "id": "183",
	  "name": "贫民窟的百万富翁 Slumdog Millionaire",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2434249040.jpg",
	  "rate": "8.6",
	  "quote": "上帝之城+猜火车+阿甘正传+开心辞典=山寨富翁",
	  "info": "导演: 丹尼·鲍尔 Danny Boyle / 洛芙琳·坦丹 Loveleen Tandan\t\t\t主演: 戴夫...",
	  "type": "2008\t/\t英国\t/\t剧情 爱情"
	 },
	 {
	  "id": "184",
	  "name": "无敌破坏王 Wreck-It Ralph",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1735642656.jpg",
	  "rate": "8.7",
	  "quote": "迪士尼和皮克斯拿错剧本的产物。",
	  "info": "导演: 瑞奇·莫尔 Rich Moore\t\t\t主演: 约翰·C·赖利 John C. Reilly / 萨拉...",
	  "type": "2012\t/\t美国\t/\t喜剧 动画 奇幻 冒险"
	 },
	 {
	  "id": "185",
	  "name": "真爱至上 Love Actually",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p475600770.jpg",
	  "rate": "8.6",
	  "quote": "爱，是个动词。",
	  "info": "导演: 理查德·柯蒂斯 Richard Curtis\t\t\t主演: 休·格兰特 Hugh Grant / 柯林...",
	  "type": "2003\t/\t英国 美国 法国\t/\t喜剧 剧情 爱情"
	 },
	 {
	  "id": "186",
	  "name": "东邪西毒 東邪西毒",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1982176012.jpg",
	  "rate": "8.6",
	  "quote": "电影诗。",
	  "info": "导演: 王家卫 Kar Wai Wong\t\t\t主演: 张国荣 Leslie Cheung / 林青霞 Brigitte...",
	  "type": "1994\t/\t中国香港 中国台湾\t/\t剧情 动作 爱情 武侠 古装"
	 },
	 {
	  "id": "187",
	  "name": "疯狂的石头",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p712241453.jpg",
	  "rate": "8.5",
	  "quote": "中国版《两杆大烟枪》。",
	  "info": "导演: 宁浩 Hao Ning\t\t\t主演: 郭涛 Tao Guo / 刘桦 Hua Liu / 连晋 Teddy Lin",
	  "type": "2006\t/\t中国大陆 中国香港\t/\t喜剧 犯罪"
	 },
	 {
	  "id": "188",
	  "name": "冰川时代 Ice Age",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1910895719.jpg",
	  "rate": "8.6",
	  "quote": "松鼠才是角儿。",
	  "info": "导演: 卡洛斯·沙尔丹哈 Carlos Saldanha / 克里斯·韦奇 Chris Wedge\t\t\t主演...",
	  "type": "2002\t/\t美国\t/\t喜剧 动画 冒险"
	 },
	 {
	  "id": "189",
	  "name": "雨中曲 Singin' in the Rain",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1612355875.jpg",
	  "rate": "9.1",
	  "quote": "骨灰级歌舞片。",
	  "info": "导演: 斯坦利·多南 Stanley Donen / 吉恩·凯利 Gene Kelly\t\t\t主演: 吉恩·...",
	  "type": "1952\t/\t美国\t/\t喜剧 歌舞 爱情"
	 },
	 {
	  "id": "190",
	  "name": "茶馆",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2528965424.jpg",
	  "rate": "9.5",
	  "info": "导演: 谢添 Tian Xie\t\t\t主演: 于是之 Shizhi Yu / 郑榕 Rong Zhen / 蓝天野 T...",
	  "type": "1982(中国大陆)\t/\t中国大陆\t/\t剧情 历史"
	 },
	 {
	  "id": "191",
	  "name": "你看起来好像很好吃 おまえうまそうだな",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p709670262.jpg",
	  "rate": "8.9",
	  "quote": "感情不分食草或者食肉。",
	  "info": "导演: 藤森雅也 Masaya Fujimori\t\t\t主演: 山口胜平 Kappei Yamaguchi / 爱河...",
	  "type": "2010\t/\t日本\t/\t剧情 动画 儿童"
	 },
	 {
	  "id": "192",
	  "name": "黑天鹅 Black Swan",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2549648344.jpg",
	  "rate": "8.6",
	  "quote": "黑暗之美。",
	  "info": "导演: 达伦·阿罗诺夫斯基 Darren Aronofsky\t\t\t主演: 娜塔莉·波特曼 Natalie...",
	  "type": "2010\t/\t美国\t/\t剧情 惊悚"
	 },
	 {
	  "id": "193",
	  "name": "恐怖游轮 Triangle",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p462470694.jpg",
	  "rate": "8.5",
	  "quote": "不要企图在重复中寻找已经失去的爱。",
	  "info": "导演: 克里斯托弗·史密斯 Christopher Smith\t\t\t主演: 梅利莎·乔治 Melissa ...",
	  "type": "2009\t/\t英国 澳大利亚\t/\t剧情 悬疑 惊悚"
	 },
	 {
	  "id": "194",
	  "name": "2001太空漫游 2001: A Space Odyssey",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2560717825.jpg",
	  "rate": "8.9",
	  "quote": "现代科幻电影的开山之作，最伟大导演的最伟大影片。",
	  "info": "导演: 斯坦利·库布里克 Stanley Kubrick\t\t\t主演: 凯尔·杜拉 Keir Dullea / ...",
	  "type": "1968\t/\t英国 美国\t/\t科幻 惊悚 冒险"
	 },
	 {
	  "id": "195",
	  "name": "魔女宅急便 魔女の宅急便",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p456676352.jpg",
	  "rate": "8.7",
	  "quote": "宫崎骏的电影总让人感觉世界是美好的，阳光明媚的。",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 高山南 Minami Takayama / 佐久间玲 Re...",
	  "type": "1989\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "196",
	  "name": "雨人 Rain Man",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2353324612.jpg",
	  "rate": "8.7",
	  "quote": "生活在自己的世界里，也可以让周围的人显得可笑和渺小。",
	  "info": "导演: 巴瑞·莱文森 Barry Levinson\t\t\t主演: 达斯汀·霍夫曼 Dustin Hoffman ...",
	  "type": "1988\t/\t美国\t/\t剧情"
	 },
	 {
	  "id": "197",
	  "name": "恋恋笔记本 The Notebook",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p483604864.jpg",
	  "rate": "8.6",
	  "quote": "爱情没有那么多借口，如果不能圆满，只能说明爱的不够。 ",
	  "info": "导演: 尼克·卡索维茨 Nick Cassavetes\t\t\t主演: 瑞恩·高斯林 Ryan Gosling /...",
	  "type": "2004\t/\t美国\t/\t剧情 爱情"
	 },
	 {
	  "id": "198",
	  "name": "遗愿清单 The Bucket List",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1481879628.jpg",
	  "rate": "8.7",
	  "quote": "用剩余不多的时间，去燃烧整个生命。",
	  "info": "导演: 罗伯·莱纳 Rob Reiner\t\t\t主演: 杰克·尼科尔森 Jack Nicholson / 摩根...",
	  "type": "2007\t/\t美国\t/\t冒险 喜剧 剧情"
	 },
	 {
	  "id": "199",
	  "name": "城市之光 City Lights",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2170238828.jpg",
	  "rate": "9.3",
	  "quote": "永远的小人物，伟大的卓别林。",
	  "info": "导演: Charles Chaplin\t\t\t主演: 查理·卓别林 Charles Chaplin / 弗吉尼亚·...",
	  "type": "1931\t/\t美国\t/\t喜剧 剧情 爱情"
	 },
	 {
	  "id": "200",
	  "name": "可可西里",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2363208684.jpg",
	  "rate": "8.9",
	  "quote": "坚硬的信仰。",
	  "info": "导演: 陆川 Chuan Lu\t\t\t主演: 多布杰 Duobujie / 张磊 Lei Zhang / 亓亮 Qi L...",
	  "type": "2004\t/\t中国大陆 中国香港\t/\t剧情 犯罪"
	 },
	 {
	  "id": "201",
	  "name": "大佛普拉斯",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2505928032.jpg",
	  "rate": "8.7",
	  "quote": "人们可以登上月球，却永远无法探索人们内心的宇宙。",
	  "info": "导演: 黄信尧 Hsin-yao Huang\t\t\t主演: 庄益增 Yizeng Zhuang / 陈竹昇 Chu-sh...",
	  "type": "2017\t/\t中国台湾\t/\t剧情 喜剧"
	 },
	 {
	  "id": "202",
	  "name": "无间道2 無間道II",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p958008320.jpg",
	  "rate": "8.7",
	  "info": "导演: 刘伟强 / 麦兆辉\t\t\t主演: 陈冠希 / 余文乐 / 曾志伟",
	  "type": "2003\t/\t中国香港\t/\t剧情 犯罪 惊悚"
	 },
	 {
	  "id": "203",
	  "name": "萤火虫之墓 火垂るの墓",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2237136718.jpg",
	  "rate": "8.7",
	  "quote": "幸福是生生不息，却难以触及的远。 ",
	  "info": "导演: 高畑勋 Isao Takahata\t\t\t主演: 辰己努 / 白石绫乃 / 志乃原良子",
	  "type": "1988\t/\t日本\t/\t动画 剧情 战争"
	 },
	 {
	  "id": "204",
	  "name": "牯岭街少年杀人事件 牯嶺街少年殺人事件",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p848381236.jpg",
	  "rate": "8.9",
	  "quote": "弱者送给弱者的一刀。",
	  "info": "导演: 杨德昌 Edward Yang\t\t\t主演: 张震 Chen Chang / 杨静怡 Lisa Yang / 张...",
	  "type": "1991\t/\t中国台湾\t/\t剧情 犯罪"
	 },
	 {
	  "id": "205",
	  "name": "背靠背，脸对脸",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2505048077.jpg",
	  "rate": "9.5",
	  "info": "导演: 黄建新 Jianxin Huang / 杨亚洲 Yazhou Yang\t\t\t主演: 牛振华 Zhenhua N...",
	  "type": "1994\t/\t中国大陆 中国香港\t/\t剧情"
	 },
	 {
	  "id": "206",
	  "name": "虎口脱险 La grande vadrouille",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2399597512.jpg",
	  "rate": "8.9",
	  "quote": "永远看不腻的喜剧。",
	  "info": "导演: 杰拉尔·乌里 Gérard Oury\t\t\t主演: 路易·德·菲耐斯 Louis de Funès...",
	  "type": "1966\t/\t法国 英国\t/\t喜剧 战争"
	 },
	 {
	  "id": "207",
	  "name": "源代码 Source Code",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p988260245.jpg",
	  "rate": "8.5",
	  "quote": "邓肯·琼斯继《月球》之后再度奉献出一部精彩绝伦的科幻佳作。",
	  "info": "导演: 邓肯·琼斯 Duncan Jones\t\t\t主演: 杰克·吉伦哈尔 Jake Gyllenhaal / ...",
	  "type": "2011\t/\t美国 加拿大\t/\t科幻 悬疑 惊悚"
	 },
	 {
	  "id": "208",
	  "name": "人工智能 Artificial Intelligence: AI",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p792257137.jpg",
	  "rate": "8.7",
	  "quote": "对爱的执着，可以超越一切。",
	  "info": "导演: 史蒂文·斯皮尔伯格 Steven Spielberg\t\t\t主演: 海利·乔·奥斯蒙 Haley...",
	  "type": "2001\t/\t美国 英国\t/\t剧情 科幻"
	 },
	 {
	  "id": "209",
	  "name": "初恋这件小事 สิ่งเล็กเล็กที่เรียกว่า...รัก",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1505312273.jpg",
	  "rate": "8.5",
	  "quote": "黑小鸭速效美白记。",
	  "info": "导演: 普特鹏·普罗萨卡·那·萨克那卡林 Puttipong Promsaka Na Sakolnakorn / 华森·波克彭...",
	  "type": "2010\t/\t泰国\t/\t剧情 喜剧 爱情"
	 },
	 {
	  "id": "210",
	  "name": "小丑 Joker",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2567198874.jpg",
	  "rate": "8.7",
	  "info": "导演: 托德·菲利普斯 Todd Phillips\t\t\t主演: 杰昆·菲尼克斯 Joaquin Phoeni...",
	  "type": "2019\t/\t美国 加拿大\t/\t剧情 犯罪 惊悚"
	 },
	 {
	  "id": "211",
	  "name": "东京教父 東京ゴッドファーザーズ",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p456703618.jpg",
	  "rate": "9.0",
	  "info": "导演: 今敏 Satoshi Kon\t\t\t主演: 江守彻 Toru Emori / 梅垣义明 Yoshiaki Ume...",
	  "type": "2003\t/\t日本\t/\t剧情 喜剧 动画"
	 },
	 {
	  "id": "212",
	  "name": "海边的曼彻斯特 Manchester by the Sea",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2421855655.jpg",
	  "rate": "8.6",
	  "quote": "我们都有权利不与自己的过去和解。",
	  "info": "导演: 肯尼斯·罗纳根 Kenneth Lonergan\t\t\t主演: 卡西·阿弗莱克 Casey Affle...",
	  "type": "2016\t/\t美国\t/\t剧情 家庭"
	 },
	 {
	  "id": "213",
	  "name": "罗生门 羅生門",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1598883511.jpg",
	  "rate": "8.8",
	  "quote": "人生的N种可能性。",
	  "info": "导演: 黑泽明 Akira Kurosawa\t\t\t主演: 三船敏郎 Toshirô Mifune / 京町子 ...",
	  "type": "1950\t/\t日本\t/\t剧情 犯罪 悬疑"
	 },
	 {
	  "id": "214",
	  "name": "终结者2：审判日 Terminator 2: Judgment Day",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1910909085.jpg",
	  "rate": "8.8",
	  "quote": "少见的超越首部的续集，动作片中的经典。",
	  "info": "导演: 詹姆斯·卡梅隆 James Cameron\t\t\t主演: 阿诺·施瓦辛格 Arnold Schwarz...",
	  "type": "1991\t/\t美国 法国\t/\t动作 科幻"
	 },
	 {
	  "id": "215",
	  "name": "青蛇",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2570901292.jpg",
	  "rate": "8.6",
	  "quote": "人生如此，浮生如斯。谁人言，花彼岸，此生情长意短。谁都是不懂爱的罢了。",
	  "info": "导演: 徐克 Hark Tsui\t\t\t主演: 张曼玉 Maggie Cheung / 王祖贤 Joey Wang / ...",
	  "type": "1993\t/\t中国香港 中国大陆\t/\t剧情 爱情 奇幻 古装"
	 },
	 {
	  "id": "216",
	  "name": "波西米亚狂想曲 Bohemian Rhapsody",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2549558913.jpg",
	  "rate": "8.6",
	  "info": "导演: 布莱恩·辛格 Bryan Singer\t\t\t主演: 拉米·马雷克 Rami Malek / 本·哈...",
	  "type": "2018\t/\t英国 美国\t/\t剧情 传记 同性 音乐"
	 },
	 {
	  "id": "217",
	  "name": "奇迹男孩 Wonder",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2507709428.jpg",
	  "rate": "8.6",
	  "quote": "世界不完美，爱会有奇迹。",
	  "info": "导演: 斯蒂芬·卓博斯基 Stephen Chbosky\t\t\t主演: 雅各布·特伦布莱 Jacob Tr...",
	  "type": "2017\t/\t美国 中国香港\t/\t剧情 儿童 家庭"
	 },
	 {
	  "id": "218",
	  "name": "二十二",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2457609817.jpg",
	  "rate": "8.7",
	  "quote": "有一些东西不应该被遗忘。",
	  "info": "导演: 郭柯 Ke Guo\t\t\t主演:",
	  "type": "2015\t/\t中国大陆\t/\t纪录片"
	 },
	 {
	  "id": "219",
	  "name": "疯狂的麦克斯4：狂暴之路 Mad Max: Fury Road",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2236181653.jpg",
	  "rate": "8.6",
	  "quote": "“多么美好的一天！”轰轰轰砰咚，啪哒哒哒轰隆隆，磅~",
	  "info": "导演: 乔治·米勒 George Miller\t\t\t主演: 汤姆·哈迪 Tom Hardy / 查理兹·塞...",
	  "type": "2015\t/\t澳大利亚 美国\t/\t动作 科幻 冒险"
	 },
	 {
	  "id": "220",
	  "name": "新龙门客栈 新龍門客棧",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1421018669.jpg",
	  "rate": "8.7",
	  "quote": "嬉笑怒骂，调风动月。",
	  "info": "导演: 李惠民 Raymond Lee\t\t\t主演: 张曼玉 Maggie Cheung / 林青霞 Brigitte ...",
	  "type": "1992\t/\t中国香港 中国大陆\t/\t动作 爱情 武侠 古装"
	 },
	 {
	  "id": "221",
	  "name": "房间 Room",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2259715855.jpg",
	  "rate": "8.8",
	  "quote": "被偷走的岁月，被伤害的生命，被禁锢的灵魂，终将被希望和善意救赎。",
	  "info": "导演: 伦尼·阿伯拉罕森 Lenny Abrahamson\t\t\t主演: 布丽·拉尔森 Brie Larson...",
	  "type": "2015\t/\t爱尔兰 加拿大 英国 美国\t/\t剧情 家庭"
	 },
	 {
	  "id": "222",
	  "name": "无耻混蛋 Inglourious Basterds",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2575043939.jpg",
	  "rate": "8.6",
	  "quote": "昆汀同学越来越变态了，比北野武还杜琪峰。",
	  "info": "导演: Quentin Tarantino\t\t\t主演: 布拉德·皮特 Brad Pitt / 梅拉尼·罗兰 M...",
	  "type": "2009\t/\t德国 美国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "223",
	  "name": "血钻 Blood Diamond",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1244017073.jpg",
	  "rate": "8.7",
	  "quote": "每个美丽事物背后都是滴血的现实。",
	  "info": "导演: 爱德华·兹威克 Edward Zwick\t\t\t主演: 莱昂纳多·迪卡普里奥 Leonardo ...",
	  "type": "2006\t/\t美国 德国 英国\t/\t剧情 惊悚 冒险"
	 },
	 {
	  "id": "224",
	  "name": "千钧一发 Gattaca",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2195672555.jpg",
	  "rate": "8.8",
	  "quote": "一部能引人思考的科幻励志片。",
	  "info": "导演: 安德鲁·尼科尔 Andrew Niccol\t\t\t主演: 伊桑·霍克 Ethan Hawke / 乌玛...",
	  "type": "1997\t/\t美国\t/\t剧情 科幻 惊悚"
	 },
	 {
	  "id": "225",
	  "name": "魂断蓝桥 Waterloo Bridge",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2351134499.jpg",
	  "rate": "8.8",
	  "quote": "中国式内在的美国电影。",
	  "info": "导演: 茂文·勒鲁瓦 Mervyn LeRoy\t\t\t主演: 费雯·丽 Vivien Leigh / 罗伯特·...",
	  "type": "1940\t/\t美国\t/\t剧情 爱情 战争"
	 },
	 {
	  "id": "226",
	  "name": "心灵奇旅 Soul",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2626308994.jpg",
	  "rate": "8.7",
	  "info": "导演: 彼特·道格特 Pete Docter / 凯普·鲍尔斯 Kemp Powers\t\t\t主演: 杰米·...",
	  "type": "2020\t/\t美国\t/\t动画 奇幻 音乐"
	 },
	 {
	  "id": "227",
	  "name": "步履不停 歩いても 歩いても",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2375245718.jpg",
	  "rate": "8.8",
	  "quote": "日本的家庭电影已经是世界巅峰了，步履不停是巅峰中的佳作。",
	  "info": "导演: 是枝裕和 Hirokazu Koreeda\t\t\t主演: 阿部宽 Hiroshi Abe / 夏川结衣 Yu...",
	  "type": "2008\t/\t日本\t/\t剧情 家庭"
	 },
	 {
	  "id": "228",
	  "name": "黑客帝国2：重装上阵 The Matrix Reloaded",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p443461390.jpg",
	  "rate": "8.7",
	  "quote": "一个精彩的世界观正在缓缓建立。",
	  "info": "导演: 拉娜·沃卓斯基 Lana Wachowski / 莉莉·沃卓斯基 Lilly Wachowski\t\t\t...",
	  "type": "2003\t/\t美国\t/\t动作 科幻"
	 },
	 {
	  "id": "229",
	  "name": "彗星来的那一夜 Coherence",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2187896734.jpg",
	  "rate": "8.5",
	  "quote": "小成本大魅力。",
	  "info": "导演: 詹姆斯·沃德·布柯特 James Ward Byrkit\t\t\t主演: 艾米丽·芭尔多尼 Em...",
	  "type": "2013\t/\t美国 英国\t/\t科幻 悬疑 惊悚"
	 },
	 {
	  "id": "230",
	  "name": "崖上的波妞 崖の上のポニョ",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2627847859.jpg",
	  "rate": "8.6",
	  "info": "导演: 宫崎骏 Hayao Miyazaki\t\t\t主演: 奈良柚莉爱 Yuria Nara / 土井洋辉 Hir...",
	  "type": "2008\t/\t日本\t/\t动画 奇幻 冒险"
	 },
	 {
	  "id": "231",
	  "name": "战争之王 Lord of War",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p792282381.jpg",
	  "rate": "8.7",
	  "quote": "做一颗让别人需要你的棋子。",
	  "info": "导演: 安德鲁·尼科尔 Andrew Niccol\t\t\t主演: 尼古拉斯·凯奇 Nicolas Cage /...",
	  "type": "2005\t/\t法国 德国 美国\t/\t剧情 犯罪"
	 },
	 {
	  "id": "232",
	  "name": "爱乐之城 La La Land",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2395517414.jpg",
	  "rate": "8.4",
	  "info": "导演: 达米恩·查泽雷 Damien Chazelle\t\t\t主演: 瑞恩·高斯林 Ryan Gosling /...",
	  "type": "2016\t/\t美国\t/\t剧情 爱情 歌舞"
	 },
	 {
	  "id": "233",
	  "name": "谍影重重2 The Bourne Supremacy",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p667644866.jpg",
	  "rate": "8.7",
	  "quote": "谁说王家卫镜头很晃？",
	  "info": "导演: 保罗·格林格拉斯 Paul Greengrass\t\t\t主演: 马特·达蒙 Matt Damon / ...",
	  "type": "2004\t/\t美国 德国\t/\t动作 悬疑 惊悚"
	 },
	 {
	  "id": "234",
	  "name": "燃情岁月 Legends of the Fall",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2797313535.jpg",
	  "rate": "8.8",
	  "quote": "传奇，不是每个人都可以拥有。",
	  "info": "导演: 爱德华·兹威克 Edward Zwick\t\t\t主演: 布拉德·皮特 Brad Pitt / 安东...",
	  "type": "1994\t/\t美国\t/\t剧情 爱情 战争 西部"
	 },
	 {
	  "id": "235",
	  "name": "阿飞正传 阿飛正傳",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2525770523.jpg",
	  "rate": "8.5",
	  "quote": "王家卫是一种风格，张国荣是一个代表。",
	  "info": "导演: 王家卫 Kar Wai Wong\t\t\t主演: 张国荣 Leslie Cheung / 张曼玉 Maggie C...",
	  "type": "1990\t/\t中国香港\t/\t犯罪 剧情 爱情"
	 },
	 {
	  "id": "236",
	  "name": "末路狂花 Thelma \u0026 Louise",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1910924635.jpg",
	  "rate": "8.9",
	  "quote": "没有了退路，只好飞向自由。",
	  "info": "导演: 雷德利·斯科特 Ridley Scott\t\t\t主演: 吉娜·戴维斯 Geena Davis / 苏...",
	  "type": "1991\t/\t美国 英国 法国\t/\t犯罪 剧情 惊悚"
	 },
	 {
	  "id": "237",
	  "name": "海洋 Océans",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2559581324.jpg",
	  "rate": "9.1",
	  "quote": "大海啊，不全是水。",
	  "info": "导演: 雅克·贝汉 Jacques Perrin / 雅克·克鲁奥德 Jacques Cluzaud\t\t\t主演:...",
	  "type": "2009\t/\t法国 瑞士 西班牙 美国 阿联酋 摩纳哥\t/\t纪录片"
	 },
	 {
	  "id": "238",
	  "name": "再次出发之纽约遇见你 Begin Again",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2250287733.jpg",
	  "rate": "8.6",
	  "quote": "爱我就给我看你的播放列表。",
	  "info": "导演: 约翰·卡尼 John Carney\t\t\t主演: 凯拉·奈特莉 Keira Knightley / 马克...",
	  "type": "2013\t/\t美国\t/\t喜剧 爱情 音乐"
	 },
	 {
	  "id": "239",
	  "name": "谍影重重 The Bourne Identity",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1597183981.jpg",
	  "rate": "8.6",
	  "quote": "哗啦啦啦啦，天在下雨，哗啦啦啦啦，云在哭泣……找自己。",
	  "info": "导演: 道格·里曼 Doug Liman\t\t\t主演: 马特·达蒙 Matt Damon / 弗兰卡·波坦...",
	  "type": "2002\t/\t美国 德国 捷克\t/\t动作 悬疑 惊悚"
	 },
	 {
	  "id": "240",
	  "name": "火星救援 The Martian",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2280097442.jpg",
	  "rate": "8.5",
	  "info": "导演: 雷德利·斯科特 Ridley Scott\t\t\t主演: 马特·达蒙 Matt Damon / 杰西卡...",
	  "type": "2015\t/\t英国 美国 匈牙利 约旦\t/\t剧情 科幻 冒险"
	 },
	 {
	  "id": "241",
	  "name": "穿越时空的少女 時をかける少女",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2079334286.jpg",
	  "rate": "8.6",
	  "quote": "爱上未来的你。 ",
	  "info": "导演: 细田守 Mamoru Hosoda\t\t\t主演: 仲里依纱 Riisa Naka / 石田卓也 Takuya...",
	  "type": "2006\t/\t日本\t/\t剧情 爱情 科幻 动画"
	 },
	 {
	  "id": "242",
	  "name": "朗读者 The Reader",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1140984198.jpg",
	  "rate": "8.6",
	  "quote": "当爱情跨越年龄的界限，它似乎能变得更久远一点，成为一种责任，一种水到渠成的相濡以沫。 ",
	  "info": "导演: 史蒂芬·戴德利 Stephen Daldry\t\t\t主演: 凯特·温丝莱特 Kate Winslet ...",
	  "type": "2008\t/\t美国 德国\t/\t剧情 爱情"
	 },
	 {
	  "id": "243",
	  "name": "香水 Perfume: The Story of a Murderer",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2571500223.jpg",
	  "rate": "8.5",
	  "quote": "一个单凭体香达到高潮的男人。",
	  "info": "导演: 汤姆·提克威 Tom Tykwer\t\t\t主演: 本·卫肖 Ben Whishaw / 艾伦·瑞克...",
	  "type": "2006\t/\t德国 法国 西班牙 美国 比利时\t/\t剧情 犯罪 奇幻"
	 },
	 {
	  "id": "244",
	  "name": "地球上的星星 Taare Zameen Par",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2197897857.jpg",
	  "rate": "8.9",
	  "quote": "天使保护事件始末。",
	  "info": "导演: 阿米尔·汗 Aamir Khan\t\t\t主演: 达席尔·萨法瑞 Darsheel Safary / 阿...",
	  "type": "2007\t/\t印度\t/\t剧情 儿童 家庭"
	 },
	 {
	  "id": "245",
	  "name": "千年女优 千年女優",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2215102596.jpg",
	  "rate": "8.8",
	  "quote": "爱情是一场没有尽头的虚幻追逐。",
	  "info": "导演: 今敏 Satoshi Kon\t\t\t主演: 庄司美代子 Miyoko Shôji / 小山茉美 Mam...",
	  "type": "2001\t/\t日本\t/\t动画 剧情 爱情"
	 },
	 {
	  "id": "246",
	  "name": "我爱你 그대를 사랑합니다",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p824590592.jpg",
	  "rate": "9.0",
	  "quote": "你要相信，这世上真的有爱存在，不管在什么年纪 ",
	  "info": "导演: 秋昌民 Chang-min Choo\t\t\t主演: 宋在河 Jae-ho Song / 李顺载 Soon-jae...",
	  "type": "2011\t/\t韩国\t/\t剧情 爱情"
	 },
	 {
	  "id": "247",
	  "name": "完美陌生人 Perfetti sconosciuti",
	  "image": "https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2522331945.jpg",
	  "rate": "8.5",
	  "quote": "来啊，互相伤害啊！",
	  "info": "导演: 保罗·格诺维瑟 Paolo Genovese\t\t\t主演: 马可·贾利尼 Marco Giallini ...",
	  "type": "2016\t/\t意大利\t/\t剧情 喜剧"
	 },
	 {
	  "id": "248",
	  "name": "驴得水",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p2393044761.jpg",
	  "rate": "8.3",
	  "quote": "过去的如果就让它过去了，未来只会越来越糟！",
	  "info": "导演: 周申 Shen Zhou / 刘露 Lu Liu\t\t\t主演: 任素汐 Suxi Ren / 大力 Da Li ...",
	  "type": "2016\t/\t中国大陆\t/\t剧情 喜剧"
	 },
	 {
	  "id": "249",
	  "name": "聚焦 Spotlight",
	  "image": "https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2263822658.jpg",
	  "rate": "8.8",
	  "quote": "新闻人的理性求真。",
	  "info": "导演: 托马斯·麦卡锡 Thomas McCarthy\t\t\t主演: 马克·鲁弗洛 Mark Ruffalo /...",
	  "type": "2015\t/\t美国\t/\t剧情 传记"
	 },
	 {
	  "id": "250",
	  "name": "浪潮 Die Welle",
	  "image": "https://img2.doubanio.com/view/photo/s_ratio_poster/public/p1344888983.jpg",
	  "rate": "8.7",
	  "quote": "世界离独裁只有五天。",
	  "info": "导演: 丹尼斯·甘塞尔 Dennis Gansel\t\t\t主演: 尤尔根·沃格尔 Jürgen Vogel ...",
	  "type": "2008\t/\t德国\t/\t剧情 惊悚"
	 }
	]```