<script>
  import Text from '../components/Text.svelte'
  import Tag from '../components/Tag.svelte'

  export let name
  export let desc
  export let link
  export let tags = []
</script>

<a href={link} class="service">
  <div class="header">
    <Text type="h2">{name}</Text>
    {#each tags as tag}
      <Tag text={tag.text} background={tag.color} />
    {/each}
  </div>
  <Text type="p" color="var(--complementary-white-50)">{desc}</Text>
</a>

<style>
  .service {
    background: var(--complementary-white-5);
    border-radius: 8px;
    padding: 16px;
    width: calc(50% - 36px);
    cursor: pointer;
    transition: 0.2s ease-in-out;
  }

  .service:hover {
    box-shadow: 0px 0px 2px 0px rgba(255, 255, 255, 0.623);
  }

  .service .header {
    display: flex;
    align-items: center;
    gap: 8px;
  }
</style>