<script>
  import Category from '../components/HubCategory.svelte'

  import Text from '../components/Text.svelte'

  export let categories = []
</script>

<div class="hub">
  <Text type="h2">Services</Text>
  <div class="content">
    {#each categories as category}
      <Category name={category.name} icon={category.icon} services={category.services} />
    {/each}
  </div>
</div>

<style>
  .hub {
    margin: 32px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .hub .content {
    display: flex;
    flex-direction: column;
    gap: 16px;
  }
</style>