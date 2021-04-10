# Mutations

### About mutations



### createTodo



#### Input fields

- input ([NewTodo!](input_objects.md#newtodo))
 

#### Returns

| Name | Description |
|------|-------------|
| done ([Boolean!](scalars.md#boolean)) |  |
| id ([ID!](scalars.md#id)) |  |
| text ([String!](scalars.md#string)) |  |
| user ([User!](objects.md#user)) |  |

---

### createUser



#### Input fields

- input ([NewUser!](input_objects.md#newuser))
 

#### Returns

| Name | Description |
|------|-------------|
| id ([ID!](scalars.md#id)) |  |
| name ([String!](scalars.md#string)) |  |
| todos ([[Todo!]!](objects.md#todo)) |  |

---

### deleteTodo



#### Input fields

- input ([DeleteTodo!](input_objects.md#deletetodo))
 

#### Returns

| Name | Description |
|------|-------------|
| done ([Boolean!](scalars.md#boolean)) |  |
| id ([ID!](scalars.md#id)) |  |
| text ([String!](scalars.md#string)) |  |
| user ([User!](objects.md#user)) |  |

---

### updateTodo



#### Input fields

- input ([UpdateTodo!](input_objects.md#updatetodo))
 

#### Returns

| Name | Description |
|------|-------------|
| done ([Boolean!](scalars.md#boolean)) |  |
| id ([ID!](scalars.md#id)) |  |
| text ([String!](scalars.md#string)) |  |
| user ([User!](objects.md#user)) |  |

---