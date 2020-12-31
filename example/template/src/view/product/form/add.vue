<style lang="less">
  @import './form.less';
</style>
<template>
  <div>


    <Form :model="formItem" :label-width="80" :rules="ruleInline">
      <FormItem label="剧本名称">
        <Input v-model="formItem.name" placeholder="剧本名称"></Input>
      </FormItem>
      <FormItem label="所得积分">
        <Input v-model="formItem.integral" placeholder="所得积分" type="number"></Input>
      </FormItem>
      <FormItem label="剧本图片">

        <Upload
          multiple
          :action="action"
          :headers="header"
          accept=".png,.PNG,.jpg,.JPG,.JPEG,.jpeg"
          :on-success="uploadSuccess"
          :on-remove="removeImg"
          :on-preview="clickImgList"
          name="img"
          :default-file-list="imgFileList"
        >
          <Button icon="ios-cloud-upload-outline" :style="{display:uploadShow}">Upload files</Button>
        </Upload>
      </FormItem>

      <FormItem>
        <Button type="primary" @click="submit">提交</Button>
        <Button style="margin-left: 8px" @click="reset">取消</Button>
      </FormItem>
    </Form>
    <Modal
      v-model="imgModal"
      title="图片预览">
      <img :src="img.remote_url" alt="" :style="{background:'url('+img.remote_url+') no-repeat top center 100%/100%', width:'450px',height:'auto'}">
    </Modal>
  </div>
</template>

<script>
  import {productAdd, productEdit, productInfo} from "@/api/data.js"
  import {getToken} from "@/libs/util";

  export default {
    data() {
      return {
        formItem: {
          name: '',
          integral: 0,
          id: 0,
          action: '',
        },
        ruleInline: {
          name: [
            {required: true, message: '请输入姓名', trigger: 'blur'}
          ],
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'}
          ],
          integral: [
            {required: true, message: '请输入积分', trigger: 'blur'},
            {type: 'number', message: '积分必须是数字', trigger: 'blur'},
          ],
        },
        type: "add",
        id: 0,
        action: "",
        header: {
          "Authorization": "Bearer " + getToken(),
        },
        img: {
          "path": "",
          "remote_url": "",
        },
        imgModal: false,
        uploadShow:"block",
        imgFileList: [],
      }
    },

    methods: {
      submit() {
        let integral = parseInt(this.formItem.integral)
        // 参数验证
        if (this.formItem.name.length == 0 || this.formItem.name.length == " ") {
          this.$Message.error("请输入剧本名称");
          return false;
        } else if (integral < 0) {
          this.$Message.error("积分不能小于0");
          return false;
        }

        // 提交
        if (this.type == "add") {
          productAdd(this.formItem.name, integral, this.img.path).then(res => {
            if (res.code != 200) {
              this.$Message.error(res.msg);
              return false;
            } else {
              this.reset();
              this.$Message.success(res.msg);
            }
          });
        } else {
          productEdit(this.formItem.id, this.formItem.name, integral, this.img.path).then(res => {
            if (res.code != 200) {
              this.$Message.error(res.msg);
              return false;
            } else {
              this.getInfo();
              this.$Message.success(res.msg);
            }
          });
        }

      },
      reset() {
        this.formItem = {
          input: '',
          mobile: '',
          integral: 0,
        }
      },
      getInfo() {
        if (this.type == "add") {
          return false;
        }
        let id = this.id
        productInfo(id).then(res => {
          if (res.code == 200) {
            this.formItem = {
              name: res.data.name,
              integral: res.data.integral,
              id: res.data.id,
            }
            // console.log(res);
            if (res.data.path.length > 0){

              let pathList=res.data.path.split("/")
              // alert(pathList[pathList.length])
              this.imgFileList = [
                {
                  "name": pathList[pathList.length-1],
                  "url": res.data.remote_url
                }
              ]
              this.img.path = res.data.image
              this.img.remote_url = res.data.remote_url
              this.uploadShow = "none"
            }

          }
        });
      },
      uploadSuccess(res, file, fileList) {
        if (res.code == 200) {
          this.uploadShow = "none"
          this.img.path = res.data.path;
          this.img.remote_url = res.data.remote_url;
        } else {
          this.$Message.error("上传失败," + res.msg);
        }
        console.log(res, file, fileList)
      },
      clickImgList() {
        this.imgModal = true;
      },
      removeImg() {
        this.img.remote_url = ""
        this.img.path = ""
        this.uploadShow = "block"
      },
    },

    mounted() {
      this.action = this.$baseUrl + "/corp/product/upload"
      // 验证是否类型
      var id = this.$route.params.id;
      if (id != undefined) {
        this.type = "edit"
        this.id = id;
        this.getInfo()
      } else {
        this.type = "add"
      }
    }
  }
</script>
