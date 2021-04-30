import trim from 'trim'
import DayjsAdapter from "@date-io/dayjs";

console.log(trim('    hello world    '))
console.log(process.env)

const dateFns = new DayjsAdapter();
const date = dateFns.date("2021/04/06");
console.log(dateFns.format(date, "fullDateTime24h"))
