<template>
  <div
    class="card new-card"
    style="width: 500px">
    <div class="card-content">
      <p class="title">{{ $t('dashboard.intents.credentials.title') }}</p>

      <b-field :label="$t('dashboard.intents.credentials.token.title')">
        <b-input
          v-model="token"
          type="text"
          rounded
          :placeholder="$t('dashboard.intents.credentials.token.text')"
          required></b-input>
      </b-field>

      <button
        class="button is-primary is-rounded"
        @click="saveToken()">
        <strong>
          {{ $t('generic.save') }}
        </strong>
      </button>
    </div>
  </div>
</template>


<script>
  export default {
    data() {
      return {
        token:  localStorage.getItem('Doctor Seuss-Token')
      }
    },
    methods: {
      saveToken() {
        localStorage.setItem('Doctor Seuss-Token', this.token)
        this.$parent.close()
      }
    }
  }
</script>
