<script>
  import Category from '../components/HubCategory.svelte'

  import Text from '../components/Text.svelte'

  let categories = [
    {
      name: "Organizing",
      icon: "flame",
      services: [
        {
          name: "Cinny",
          desc: "Chat safely using matrix",
          link: "https://organize.coldwire.org"
        },
        {
          name: "Bloc",
          desc: "Securely store your files",
          tags: [
            {text: "soon", color: "green"}
          ]
        },
      ]
    },
    {
      name: "Hosting",
      icon: "hosting",
      services: [
        {
          name: "VPS",
          desc: "Run your own server anonymously",
          tags: [
            {text: "soon", color: "green"},
            {text: "On request", color: "purple"}
          ]
        },
      ]
    }
  ]
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
    margin-top: 92px;
    margin-left: 32px;
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