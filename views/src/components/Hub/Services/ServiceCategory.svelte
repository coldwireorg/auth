<script>
  import Item from './ServicesItem.svelte';
  import Text from '../../Text.svelte'

  export let name
  export let icon
  export let services = []
</script>

<div class="category">
  <div class="header">
    <img src="/icons/{icon}.svg" alt="">
    <Text type="h3">{name}</Text>
  </div>

  <div class="services">
    {#each services as service}
      <Item name={service.name} desc={service.desc} tags={service.tags} link={service.link} />
    {/each}
  </div>
</div>

<style>
  .category {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding: 16px;
    gap: 16px;

    background: var(--complementary-gray-3);
    border-radius: 8px;

    max-width: 496px;
  }

  .category .header {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .category .header img {
    height: 16px;
  }

  .category .services {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    width: 100%;
  }
</style>