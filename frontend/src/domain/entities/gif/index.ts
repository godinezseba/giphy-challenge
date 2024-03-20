export type GIF = {
  name: string
  URL: string
  content: string
  user: string
  tags: string[]
  createdAt: Date
};

export const makeGIF = (
  input: Partial<GIF>,
): GIF => {
  let name: string = input.name ?? '';
  let URL: string = input.URL ?? '';
  let content: string = input.content ?? '';
  let user: string = input.user ?? '';
  let tags: string[] = input.tags ?? [];
  let createdAt: Date = input.createdAt ?? new Date();

  return {
    get name(): string {
      return name;
    },
    set name(newName: string) {
      name = newName;
    },
    get URL(): string {
      return URL;
    },
    set URL(newURL: string) {
      URL = newURL;
    },
    get content(): string {
      return content;
    },
    set content(newContent: string) {
      content = newContent;
    },
    get user(): string {
      return user;
    },
    set user(newUser: string) {
      user = newUser;
    },
    get tags(): string[] {
      return tags;
    },
    set tags(newTags: string[]) {
      tags = newTags;
    },
    get createdAt(): Date {
      return createdAt;
    },
    set createdAt(newCreatedAt: Date) {
      createdAt = newCreatedAt;
    },
  };
};
