import { h, ref } from 'vue';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { defRangeShortcuts } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import {DictOptions, getComputedOptions, renderOptionTagWithOpts} from '@/utils';
export class State {
  public id = 0; // 用户ID
  public pid = 0; // 父级ID
  public comId = 0; // 小区ID
  public houseId = 0; // 房屋ID
  public username = ''; // 姓名
  public mobile = ''; // 手机号
  public password = ''; // 密码
  public passwordSalt = ''; // salt
  public idCard = ''; // 身份证
  public isAuthed = -1; // 是否已认证
  public authedAt = ''; // 最新认证时间
  public totalParkTime = 0; // 累计租住时长
  public userType = 1; // 租户类型
  public userStatus = 1; // 用户状态
  public startAt = ''; // 开始日期
  public endAt = ''; // 结束日期
  public remark = ''; // 备注
  public nowIsPay = -1; // 本期是否已支付
  public idCardFront = ''; // 身份证正面
  public idCardBack = ''; // 身份证背面
  public createdAt = ''; // 创建时间
  public updatedAt = ''; // 修改时间
  public refoundRs = ''; // 退租理由
  public isChangedPass = -1; // 是否已更换密码
  public openId = ''; // openId

  constructor(state?: Partial<State>) {
    if (state) {
      Object.assign(this, state);
    }
  }
}
export function newState(state: State | Record<string, any> | null): State {
  if (state !== null) {
    if (state instanceof State) {
      return cloneDeep(state);
    }
    return new State(state);
  }
  return new State();
}

// todo
const optMap = ref<DictOptions>({})
export function setOptMap(optMp: any) {
  optMap.value = optMp;
}

// 表单验证规则
export const rules = {
  comId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入小区ID',
  },
  houseId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入房屋ID',
  },
  username: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入姓名',
  },
  mobile: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.phone,
  },
  password: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.password,
  },
  passwordSalt: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入salt',
  },
  idCard: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.idCard,
  },
  isAuthed: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否已认证',
  },
  totalParkTime: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入累计租住时长',
  },
  userType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入租户类型',
  },
  userStatus: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入用户状态',
  },
  nowIsPay: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入本期是否已支付',
  },
  idCardFront: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入身份证正面',
  },
  idCardBack: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入身份证背面',
  },
  isChangedPass: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入是否已更换密码',
  },
};

// 表格搜索表单
export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '用户ID',
    componentProps: {
      placeholder: '请输入用户ID',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'userStatus',
    component: 'NSelect',
    label: '用户状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择用户状态',
      options: getComputedOptions('sys_login_status',optMap), //todo
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

// 表格列
export const columns = [
  {
    title: '用户ID',
    key: 'id',
    align: 'left',
    width: -1,
  },
  {
    title: '父级ID',
    key: 'pid',
    align: 'left',
    width: -1,
  },
  {
    title: '小区ID',
    key: 'comId',
    align: 'left',
    width: -1,
  },
  {
    title: '房屋ID',
    key: 'houseId',
    align: 'left',
    width: -1,
  },
  {
    title: '姓名',
    key: 'username',
    align: 'left',
    width: -1,
  },
  {
    title: '手机号',
    key: 'mobile',
    align: 'left',
    width: -1,
  },
  {
    title: '密码',
    key: 'password',
    align: 'left',
    width: -1,
  },
  {
    title: 'salt',
    key: 'passwordSalt',
    align: 'left',
    width: -1,
  },
  {
    title: '身份证',
    key: 'idCard',
    align: 'left',
    width: -1,
  },
  {
    title: '是否已认证',
    key: 'isAuthed',
    align: 'left',
    width: -1,
  },
  {
    title: '最新认证时间',
    key: 'authedAt',
    align: 'left',
    width: -1,
  },
  {
    title: '累计租住时长',
    key: 'totalParkTime',
    align: 'left',
    width: -1,
  },
  {
    title: '租户类型',
    key: 'userType',
    align: 'left',
    width: -1,
  },
  {
    title: '用户状态',
    key: 'userStatus',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTagWithOpts('sys_login_status', row.userStatus,optMap.value); // todo
    },
  },
  {
    title: '开始日期',
    key: 'startAt',
    align: 'left',
    width: -1,
  },
  {
    title: '结束日期',
    key: 'endAt',
    align: 'left',
    width: -1,
  },
  {
    title: '备注',
    key: 'remark',
    align: 'left',
    width: -1,
  },
  {
    title: '本期是否已支付',
    key: 'nowIsPay',
    align: 'left',
    width: -1,
  },
  {
    title: '创建时间',
    key: 'createdAt',
    align: 'left',
    width: -1,
  },
  {
    title: '修改时间',
    key: 'updatedAt',
    align: 'left',
    width: -1,
  },
  {
    title: '退租理由',
    key: 'refoundRs',
    align: 'left',
    width: -1,
  },
  {
    title: '是否已更换密码',
    key: 'isChangedPass',
    align: 'left',
    width: -1,
    render(row: State) {
      return renderOptionTagWithOpts('sys_normal_disable', row.isChangedPass,optMap.value); // todo
    },
  },
  {
    title: 'openId',
    key: 'openId',
    align: 'left',
    width: -1,
  },
];
