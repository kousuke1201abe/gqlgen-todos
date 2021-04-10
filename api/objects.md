# Objects

### About objects

[Objects](https://graphql.github.io/graphql-spec/June2018/#sec-Objects) in GraphQL represent the resources you can access. An object can contain a list of fields, which are specifically typed.

### Todo

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>done</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>text</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>user</strong> (<a href="objects.md#user">User!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### User

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>todos</strong> (<a href="objects.md#todo">[Todo!]!</a>)</td> 
    <td></td>
  </tr>
</table>

---