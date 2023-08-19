/* A very simple mail database */
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SIZE 100

struct list_type {
    char name[40];
    char street[40];
    char city[40];
    char state[3];
    char zip[10];
} list[SIZE];

/* Function prototypes */
char menu(void);
void init_list(void), enter(void), display(void), save(void), load(void),
    clear(void);

int main(void) {
    char choice;
    init_list();

    for (;;) {
        choice = menu();
        switch (choice) {
        case 'e':
            enter();
            break;
        case 'd':
            display();
            break;
        case 's':
            save();
            break;
        case 'l':
            load();
            break;
        case 'c':
            clear();
            break;
        case 'q':
            exit(0);
        }
    }
}

/* init list */
void init_list() {
    register int t;
    for (t = 0; t < SIZE; t++)
        *list[t].name = '\0';
    /* zero length name indicates empty */
    printf("Done Initializing list...\n");
}

void enter() {
    register int i;
    for (i = 0; i < SIZE; i++) {
        if (*list[i].name == '\0')
            break;
    }

    if (i == SIZE) {
        printf("List FUll");
        return;
    }

    printf("Name: ");
    fgets(list[i].name, sizeof(list[i].name), stdin);

    printf("Street: ");
    fgets(list[i].street, sizeof(list[i].name), stdin);

    printf("City: ");
    fgets(list[i].city, sizeof(list[i].name), stdin);

    printf("State: ");
    fgets(list[i].state, sizeof(list[i].name), stdin);

    printf("ZIP: ");
    fgets(list[i].zip, sizeof(list[i].name), stdin);
}

void display(void) {
    register int t;
    for (t = 0; t < SIZE; t++) {
        if (*list[t].name) {
            printf("%s", list[t].name);
            printf("%s", list[t].street);
            printf("%s", list[t].city);
            printf("%s", list[t].state);
            printf("%s", list[t].zip);
        }
    }
}
void save(void) { // save the list.
    FILE *fp;
    register int t;
    fp = fopen("maillist", "ab");
    if (fp == NULL) {
        printf("Error Opening Records");
        return;
    }

    for (t = 0; t < SIZE; t++) {
        if (*list[t].name) {
            if (fwrite(&list[t], sizeof(struct list_type), 1, fp) != 1) {
                printf("Error Writing Record");
            }
        }
    }

    fclose(fp);
}
void load(void) { // load the list
    FILE *fp;
    register int t;
    fp = fopen("maillist", "rb");
    if (fp == NULL) {
        printf("Error Opening Records");
        return;
    }

    init_list();
    for (t = 0; t < SIZE; t++) {
        if (fread(&list[t], sizeof(struct list_type), 1, fp) != 1) {
            printf("File Read Error");
            return;
        }
    }

    fclose(fp);
}

void clear(void) {
    FILE *fp;
    fp = fopen("maillist", "wb");
    if (fp == NULL) {
        printf("Error Opening Records");
        return;
    }
    fclose(fp);
    printf("Database has been cleared \n");
}

char menu(void) { // menu
    char s[80];
    do {
        printf("\n");
        printf("(E)nter \n");
        printf("(D)isplay \n");
        printf("(L)oad \n");
        printf("(S)ave \n");
        printf("(C)lear \n");
        printf("(Q)uit \n");
        printf("\n");
        printf("Pick an Option: ");
        scanf("%s", s);
        printf("\n");

        int c;
        while ((c = getchar()) != '\n' && c != EOF)
            ; // clear the input buffer
    } while (!strchr("edlscq", tolower(*s)));
    return tolower(*s);
}
