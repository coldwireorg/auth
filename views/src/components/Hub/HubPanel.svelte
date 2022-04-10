<script>
  import Text from '../../components/Text.svelte'

  export let title = ""
</script>

<div class="panel">
  <Text type="h2">{title}</Text>
  <div class="content">
    <slot/>
  </div>
</div>

<style>
  .panel {
    display: flex;
    flex-direction: column;
    gap: 16px;    
  }

  .panel .content {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }
</style>