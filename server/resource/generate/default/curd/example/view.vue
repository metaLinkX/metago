<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="租户详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                父级ID
              </template>
              {{ formValue.pid }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                小区ID
              </template>
              {{ formValue.comId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                房屋ID
              </template>
              {{ formValue.houseId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                姓名
              </template>
              {{ formValue.username }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                手机号
              </template>
              {{ formValue.mobile }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                密码
              </template>
              {{ formValue.password }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                salt
              </template>
              {{ formValue.passwordSalt }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                身份证
              </template>
              {{ formValue.idCard }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否已认证
              </template>
              {{ formValue.isAuthed }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最新认证时间
              </template>
              {{ formValue.authedAt }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                累计租住时长
              </template>
              {{ formValue.totalParkTime }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                租户类型
              </template>
              {{ formValue.userType }}
            </n-descriptions-item>
            <n-descriptions-item label="用户状态">
              <n-tag :type="getDictType('sys_login_status', formValue.userStatus)" size="small" class="min-left-space">
                {{ getDictLabel('sys_login_status', formValue.userStatus) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                开始日期
              </template>
              {{ formValue.startAt }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                结束日期
              </template>
              {{ formValue.endAt }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                备注
              </template>
              {{ formValue.remark }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                本期是否已支付
              </template>
              {{ formValue.nowIsPay }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                身份证正面
              </template>
              <span v-html="formValue.idCardFront"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                身份证背面
              </template>
              <span v-html="formValue.idCardBack"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                退租理由
              </template>
              {{ formValue.refoundRs }}
            </n-descriptions-item>
            <n-descriptions-item label="是否已更换密码">
              <n-tag :type="getDictType('sys_normal_disable', formValue.isChangedPass)" size="small" class="min-left-space">
                {{ getDictLabel('sys_normal_disable', formValue.isChangedPass) }}
              </n-tag>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                openId
              </template>
              {{ formValue.openId }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { View } from '@/api/hgUser';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });


  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>
