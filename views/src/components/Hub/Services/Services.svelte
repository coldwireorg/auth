<script>
  import Category from './ServiceCategory.svelte'
  import Text from '../../Text.svelte'

  let categories = [
    {
      name: "Organizing",
      icon: "flame",
      services: [
        {
          name: "Element",
          desc: "Chat safely using matrix",
          link: "https://coldwire.org/_matrix/client/r0/login/sso/redirect?redirectUrl=https://organize.coldwire.org"
        },
        {
          name: "Cinny",
          desc: "Simple matrix client",
          link: "https://coldwire.org/_matrix/client/r0/login/sso/redirect?redirectUrl=https://cinny.coldwire.org"
        },
        {
          name: "Bloc",
          desc: "Securely store your files",
          tags: [
            {text: "soon", color: "green"}
          ]
        },
        {
          name: "Form",
          desc: "Encrypted and anonymous surveys",
          tags: [
            {text: "soon", color: "green"}
          ]
        },
        {
          name: "Riot",
          desc: "Organize protests",
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
        {
          name: "Gitea",
          desc: "Share your code and collaborate",
          tags: [
            {text: "soon", color: "green"},
          ]
        },
      ]
    },
    {
      name: "Privacy",
      icon: "eye",
      services: [
        {
          name: "Reflect",
          desc: "Decentralized onion routed VPN",
          tags: [
            {text: "soon", color: "green"},
          ]
        },
      ]
    }
  ]
</script>

<div class="content">
  {#each categories as category}
    <Category name={category.name} icon={category.icon} services={category.services} />
  {/each}
</div>

<style>
  .content {
    display: flex;
    flex-direction: column;
    gap: 16px;
    overflow: auto;
  }
</style>