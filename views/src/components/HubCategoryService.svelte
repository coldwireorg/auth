<script>
  import Text from '../components/Text.svelte'
  import Tag from '../components/Tag.svelte'

  export let name
  export let desc
  export let tags = []
</script>

<div class="service">
  <div class="header">
    <Text type="h2">{name}</Text>
    {#each tags as tag}
      <Tag text={tag.text} background={tag.color} />
    {/each}
  </div>
  <Text type="p" color="var(--complementary-white-50)">{desc}</Text>
</div>

<style>
  .service {
    background: #30363D;
    border-radius: 8px;
    padding: 16px;
    width: calc(50% - 36px);
  }

  .service .header {
    display: flex;
    align-items: center;
    gap: 8px;
  }
</style>