USE course_enrollment_manager;

CREATE TABLE students(
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE teachers(
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL
);

CREATE TABLE courses(
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    semester INT NOT NUll, -- 1: first semester　2:second semester 3:full year 4:others
    instructor INT,
    credits INT NOT NUll,
    description TEXT,
    FOREIGN KEY (instructor) REFERENCES teachers(id)
);

CREATE TABLE students_courses(
    student_id INT ,
    course_id INT,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE course_schedules (
    id INT PRIMARY KEY AUTO_INCREMENT,
    course_id INT,
    day_of_week INT,
    period INT,
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- init data
INSERT INTO students (`name`,`password`,email) VALUES
("鈴木太郎","$2a$10$Wd3u8SEBuT4UQ3IyEOFX5OHZjcgV8N9SrnR9N2mhyZcJNylHhfAXO","suzuki@scon.com");

INSERT INTO teachers (`name`,position) VALUES
("飯田涼","准教授"),
("佐々木比奈","教授"),
("Aaron Swift","講師"),
("佐藤武","講師"),
("田中孝司","教授"),
("秋田達","講師"),
("WangXin","講師");

INSERT INTO courses (`name`,semester,instructor,credits,description) VALUES
("経済学(上)",1,1,2,
"現実の経済問題を豊富に取り上げ、初学者が経済学を具体的に理解できるよう工夫。\n
Part １前期　ミクロ経済学\n
　１　需要と供給\n
　２　需要曲線と消費者行動\n
　３　費用の構造と供給行動\n
　４　市場取引と資源配分\n
　５　独占と競争の理論\n
　６　市場の失敗\n
　７　不確実性と不完全情報の世界\n
　８　ゲームの理論入門\n
\n
Part ２後期　マクロ経済学\n
　９　経済をマクロからとらえる\n
　10　有効需要と乗数メカニズム\n
　11　貨幣の機能\n
　12　マクロ経済政策\n
　13　インフレ・デフレと失業\n
　14　高齢社会の財政運営\n
　15　経済成長と経済発展\n
　16　国際経済学
"),
("経済学(下)",2,1,2,
"現実の経済問題を豊富に取り上げ、初学者が経済学を具体的に理解できるよう工夫。\n
Part １前期　ミクロ経済学\n
　１　需要と供給\n
　２　需要曲線と消費者行動\n
　３　費用の構造と供給行動\n
　４　市場取引と資源配分\n
　５　独占と競争の理論\n
　６　市場の失敗\n
　７　不確実性と不完全情報の世界\n
　８　ゲームの理論入門\n
\n
Part ２後期　マクロ経済学\n
　９　経済をマクロからとらえる\n
　10　有効需要と乗数メカニズム\n
　11　貨幣の機能\n
　12　マクロ経済政策\n
　13　インフレ・デフレと失業\n
　14　高齢社会の財政運営\n
　15　経済成長と経済発展\n
　16　国際経済学
"),
("公共経済学理論",3,2,2,
"自由競争市場を基本とする経済社会において，公共部門の果たす役割を研究対象とする学問。 市場の失敗を前提とし，政策的にどのような解決方法があるか分析することを主な目的とする。\n
シバラス\n
第１部 公共部門の役割：イントロダクション \n
第１章 公共部門とは何か\n
第２部 厚生経済学の基礎\n
第２章 市場の効率性\n
第３章 市場の失敗\n
第４章 公共財と公的に供給される私的財\n
第５章 外部性と環境問題\n
第６章 効率と公平\n
第３部 公共支出の理論\n
第７章 財・サービスの公的生産．\n
第８章 公共選択\n
第４部 支出計画\n
第９章 支出政策の分析枠組み\n
第１０章 公共支出の評価\n
第１１章 国防・研究・技術\n
第１２章 医療・教育\n
第１３章 社会保障制度と所得再分配\n
授業単位：2
"),
("Information Economics",1,3,4,
"Information economics or the economics of information is the branch of microeconomics that studies how information and information systems affect an economy and economic decisions.\n

One application considers information embodied in certain types of commodities that are expensive to produce but cheap to reproduce. \n
Examples include computer software (e.g., Microsoft Windows), pharmaceuticals and technical books. Without the basic research, initial production of high-information commodities may be too unprofitable to market, a type of market failure. Government subsidization of basic research has been suggested as a way to mitigate the problem.\n
Contents\n
Part 1: Information as an Economic Good 3. What is Information? 4. The Value of Information 5. The Optimal Amount of Information 6. The Production of Information\n
Part 2: How the Market Aggregates Information 7. From Information to Prices 8. Knowing Facts or Reading Thoughts 9. Coordination Problems 10. Learning and Cascades 11. The Macroeconomics of Information \n
Part 3: The Economics of Information Asymmetries 12. The winner's Curse 13. Hidden Information and Self-Selection 14. Optimal Contracts 15. The Revelation Principle 16. Creating Incentives "
),
("英語",4,1,2,
"レベルB2～C1（B2 社会生活での幅広い話題について自然に会話ができ、明確かつ詳細に自分の意見を表現できる／C1 広範で複雑な話題を理解して、目的に合った適切な言葉を使い、論理的な主張や議論を組み立てることができる）\n
教科書：英語長文問題精講＆American Pie Slice of Life Essays on America and Japan \n
授業単位：2
");

INSERT INTO course_schedules (course_id,day_of_week,period) VALUES
(1,3,2),(2,3,4),(3,1,4),(4,5,3),(4,5,4),(5,2,3),(5,4,3);

INSERT INTO students_courses (student_id,course_id) VALUES
(1,2),(1,3);

INSERT INTO courses (`name`,semester,instructor,credits,description) VALUES
("中級途上国経済　分野別演習",1,4,2,
"現中級途上国経済に取り上げ、中級途上国の経済を具体的に理解できるよう工夫。\n
「中級途上国経済」での学習内容に沿いながら、発展途上国が直面する開発課題について、具体的な国や地域を決めて学習と研究を進め、課題レポートを作成していきます。\n
労働問題、住宅問題、財政問題、技術革新、貨幣・通貨などトピックは各自が自由に選び、英語文献やできれば現地語文献を用いてサーベイを進めていきます。
新しい研究課題を発見し、かつその課題の解決方法のための方向性や可能性を見出すことが演習の目標です。
"),
("ミクロ経済学入門",3,5,4,
"ミクロ経済学入門
"),
("体育",3,5,2,
"体育
"),
("共通実践",3,6,4,
"共通実践Ⅰ
"),
("社会学",1,6,2,
"社会学
"),
("中国語入門（１） ",1,7,2,
"中国語の基本を身につけるビギナーにぴったりの1冊です。\n
初学者に必須の「発音」と「文法」の両方をマスターできます。\n
「発音」→「文法」という構成で、まず「発音編」で声調、子音・母音、ピンインを身につけて、「文法編」で基礎文法を段階的に積み上げていきます。\n
文法の説明は45の「文法公式」にまとめて紹介しているので、頭を整理しながらシンプルに覚えていくことができます。\n
文法公式の後には、書き込むスタイルのエクササイズが用意されていて、文法知識を定着させながら、「簡体字」を覚えられます。\n"),
("中国語入門（２））",2,7,2,
"中国語の基本を身につけるビギナーにぴったりの1冊です。\n
初学者に必須の「発音」と「文法」の両方をマスターできます。\n
「発音」→「文法」という構成で、まず「発音編」で声調、子音・母音、ピンインを身につけて、「文法編」で基礎文法を段階的に積み上げていきます。\n
文法の説明は45の「文法公式」にまとめて紹介しているので、頭を整理しながらシンプルに覚えていくことができます。\n
文法公式の後には、書き込むスタイルのエクササイズが用意されていて、文法知識を定着させながら、「簡体字」を覚えられます。\n
");

INSERT INTO course_schedules (course_id,day_of_week,period) VALUES
(6,1,2),(7,1,3),(8,4,2),(9,4,3),(10,3,1),(11,2,1),(12,4,3);