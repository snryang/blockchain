/*辅助库类 */

let _firstTo = p => R.compose(
    R.join(''),
    R.juxt([R.compose(p, R.head), R.tail])
);

/* yang => Yang */
let firstToUpper = firstTo(R.toUpper);
/* YANG => yANG */
let firstToLower = firstTo(R.toLower);
/* yang_wen_bing => yangWenBing */
let camelCase=R.compose(firstToLower,R.join(''),R.map(firstToUpper), R.split("_"))
