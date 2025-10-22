<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑租户 #' + formValue.id : '添加租户'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="父级ID" path="pid">
                  <n-input-number placeholder="请输入父级ID" v-model:value="formValue.pid" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="小区ID" path="comId">
                  <n-input-number placeholder="请输入小区ID" v-model:value="formValue.comId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="房屋ID" path="houseId">
                  <n-input-number placeholder="请输入房屋ID" v-model:value="formValue.houseId" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="姓名" path="username">
                  <n-input placeholder="请输入姓名" v-model:value="formValue.username" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手机号" path="mobile">
                  <n-input placeholder="请输入手机号" v-model:value="formValue.mobile" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="密码" path="password">
                  <n-input placeholder="请输入密码" v-model:value="formValue.password" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="salt" path="passwordSalt">
                  <n-input placeholder="请输入salt" v-model:value="formValue.passwordSalt" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="身份证" path="idCard">
                  <n-input placeholder="请输入身份证" v-model:value="formValue.idCard" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否已认证" path="isAuthed">
                  <n-input-number placeholder="请输入是否已认证" v-model:value="formValue.isAuthed" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最新认证时间" path="authedAt">
                  <DatePicker v-model:formValue="formValue.authedAt" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="累计租住时长" path="totalParkTime">
                  <n-input-number placeholder="请输入累计租住时长" v-model:value="formValue.totalParkTime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="租户类型" path="userType">
                  <n-input-number placeholder="请输入租户类型" v-model:value="formValue.userType" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="用户状态" path="userStatus">
                  <n-select v-model:value="formValue.userStatus" :options="sys_login_statusDictOptions" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="开始日期" path="startAt">
                  <DatePicker v-model:formValue="formValue.startAt" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="结束日期" path="endAt">
                  <DatePicker v-model:formValue="formValue.endAt" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="备注" path="remark">
                  <n-input placeholder="请输入备注" v-model:value="formValue.remark" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="本期是否已支付" path="nowIsPay">
                  <n-input-number placeholder="请输入本期是否已支付" v-model:value="formValue.nowIsPay" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="身份证正面" path="idCardFront">
                  <Editor style="height: 450px" id="idCardFront" v-model:value="formValue.idCardFront" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="身份证背面" path="idCardBack">
                  <Editor style="height: 450px" id="idCardBack" v-model:value="formValue.idCardBack" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退租理由" path="refoundRs">
                  <n-input placeholder="请输入退租理由" v-model:value="formValue.refoundRs" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否已更换密码" path="isChangedPass">
                  <n-select v-model:value="formValue.isChangedPass" :options="sys_normal_disableDictOptions" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="openId" path="openId">
                  <n-input placeholder="请输入openId" v-model:value="formValue.openId" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { useDictStore } from '@/store/modules/dict';
  import { Edit, View } from '@/api/hgUser';
  import { State, newState, rules } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import Editor from '@/components/Editor/editor.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();


  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;

    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);

      return;
    }

    // 编辑
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

<style lang="less"></style>
