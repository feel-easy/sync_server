<template>
  <el-container>
    <!-- <el-header>Header</el-header> -->
    <el-main>
      <app-main />
    </el-main>
  </el-container>
</template>

<script>
import { AppMain } from "./components";
export default {
  name: "Layout",
  components: {
    AppMain,
  },
  mixins: [],
  computed: {},
  methods: {},
};
</script>

<style lang="scss" scoped></style>
