
int puissance(int a, int b) {
	int r;
    r = 1;
	while(b!=0){
		r = r*a;
		b = b-1;
	}
	return 1;
}

int free(int p){
	return 0;
}

int printsub (int n){
	if(n==0){
		return 0;
	}
	int r;
	int d;
	
	r = n/10;
	d = n%10;
	
	printsub(r);
	send d+48;
	return 0;
}

int print(int n){
	if (n<0){
		send 45;
		n = -n;
	}
	if (n==0){
		send 48;
	}else{
		printsub(n);
	}
}

int malloc(int n){
	int p;
    p = *0;
	
	*0 = *0 + n;
	return p;
}



